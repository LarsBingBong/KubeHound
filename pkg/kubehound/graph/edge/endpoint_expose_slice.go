package edge

import (
	"context"
	"fmt"

	"github.com/DataDog/KubeHound/pkg/kubehound/graph/adapter"
	"github.com/DataDog/KubeHound/pkg/kubehound/graph/types"
	"github.com/DataDog/KubeHound/pkg/kubehound/models/converter"
	"github.com/DataDog/KubeHound/pkg/kubehound/storage/cache"
	"github.com/DataDog/KubeHound/pkg/kubehound/storage/storedb"
	"github.com/DataDog/KubeHound/pkg/kubehound/store/collections"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func init() {
	Register(&EndpointExposeSlice{}, RegisterDefault)
}

type EndpointExposeSlice struct {
	BaseEdge
}

type sliceEndpoitGroup struct {
	Endpoint  primitive.ObjectID `bson:"_id" json:"endpoint_id"`
	Container primitive.ObjectID `bson:"container_id" json:"container_id"`
}

func (e *EndpointExposeSlice) Label() string {
	return "ENDPOINT_EXPOSE"
}

func (e *EndpointExposeSlice) Name() string {
	return "EndpointExposeSlice"
}

func (e *EndpointExposeSlice) Processor(ctx context.Context, oic *converter.ObjectIDConverter, entry any) (any, error) {
	typed, ok := entry.(*sliceEndpoitGroup)
	if !ok {
		return nil, fmt.Errorf("invalid type passed to processor: %T", entry)
	}

	return adapter.GremlinEdgeProcessor(ctx, oic, e.Label(), typed.Endpoint, typed.Container)
}

func (e *EndpointExposeSlice) Stream(ctx context.Context, store storedb.Provider, c cache.CacheReader,
	callback types.ProcessEntryCallback, complete types.CompleteQueryCallback) error {

	endpoints := adapter.MongoDB(store).Collection(collections.EndpointName)

	// K8s endpoint slices must be ingested before containers. In this stage we need to match store.Endpoint documents that
	// are generated via K8s EndpointSlice objects and match them to the container exposing the endpoint. The other case of
	// store.Endpoint documents not associated with an EndpointSlice is handled separately.
	pipeline := []bson.M{
		{
			// Match only endpoints with a matching EndpointSlice
			"$match": bson.M{
				"has_slice": true,
			},
		},
		{
			// Lookup the container matching the slice. This requires a match on namespace/pod/port/protocol
			"$lookup": bson.M{
				"as":   "matchContainers",
				"from": "containers",
				"let": bson.M{
					"pod":   "$pod_name",
					"podNS": "$pod_namespace",
					"port":  "$port.port",
					"proto": "$port.protocol",
				},
				"pipeline": []bson.M{
					{
						"$match": bson.M{"$expr": bson.M{
							"$and": bson.A{
								bson.M{"$eq": bson.A{
									"$inherited.namespace", "$$podNS",
								}},
								bson.M{"$eq": bson.A{
									"$inherited.pod_name", "$$pod",
								}},
								bson.M{"$ne": bson.A{
									"$k8.ports", nil,
								}},
								// Cannot use an $elemMatch with pipeline variables so use the more convoluted $filter syntax to match container port/protocol
								// See: https://www.mongodb.com/community/forums/t/equivalent-of-elemmatch-query-operator-for-use-in-match-within-the-aggregation-lookup-with-pipeline/5360
								bson.M{"$gt": bson.A{
									bson.M{"$size": bson.M{"$filter": bson.M{
										"input": "$k8.ports",
										"as":    "p",
										"cond": bson.M{
											"$and": bson.A{
												bson.M{"$eq": bson.A{
													"$$p.containerport", "$$port",
												}},
												bson.M{"$eq": bson.A{
													"$$p.protocol", "$$proto",
												}},
											}},
									}}},
									0,
								}},
							},
						}},
					},
					{
						"$project": bson.M{
							"_id": 1,
						},
					},
				},
			},
		},
		{
			"$unwind": "$matchContainers",
		},
		{
			"$project": bson.M{
				"_id":          1,
				"container_id": "$matchContainers._id",
			},
		},
	}

	cur, err := endpoints.Aggregate(context.Background(), pipeline)
	if err != nil {
		return err
	}
	defer cur.Close(ctx)

	return adapter.MongoCursorHandler[sliceEndpoitGroup](ctx, cur, callback, complete)
}
