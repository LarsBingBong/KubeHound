package system

import (
	"context"
	"testing"

	"github.com/DataDog/KubeHound/pkg/config"
	"github.com/DataDog/KubeHound/pkg/kubehound/graph/vertex"
	"github.com/DataDog/KubeHound/pkg/kubehound/models/graph"
	"github.com/DataDog/KubeHound/pkg/kubehound/models/shared"
	"github.com/DataDog/KubeHound/pkg/kubehound/storage/graphdb"
	gremlingo "github.com/apache/tinkerpop/gremlin-go/driver"
	"github.com/stretchr/testify/suite"
)

//go:generate go run ./generator ../setup/test-cluster ./vertex.gen.go

// Optional syntactic sugar.
var __ = gremlingo.T__

type VertexTestSuite struct {
	suite.Suite
	gdb    graphdb.Provider
	client *gremlingo.DriverRemoteConnection
}

func (suite *VertexTestSuite) SetupTest() {
	gdb, err := graphdb.Factory(context.Background(), config.MustLoadConfig("./kubehound.yaml"))
	suite.NoError(err)
	suite.gdb = gdb
	suite.client = gdb.Raw().(*gremlingo.DriverRemoteConnection)
	// g := gremlingo.Traversal_().WithRemote(suite.client)
	// errChan := g.V().Drop().Iterate()
	// err = <-errChan
	// if err != nil {
	// 	suite.Errorf(err, "error deleting the graphdb:\n")
	// }
	// err = runKubeHound()
	// suite.NoError(err)
}

func (suite *VertexTestSuite) TestVertexContainer() {
	g := gremlingo.Traversal_().WithRemote(suite.client)
	results, err := g.V().HasLabel(vertex.ContainerLabel).ElementMap().ToList()
	suite.NoError(err)

	suite.Equal(len(expectedContainers), len(results))
	resultsMap := map[string]graph.Container{}
	for _, res := range results {
		res := res.GetInterface()
		converted := res.(map[any]any)

		nodeName, ok := converted["name"].(string)
		suite.True(ok, "failed to convert node name to string")

		compromised, ok := converted["compromised"].(int)
		suite.True(ok, "failed to convert compromised field to CompromiseType")

		critical, ok := converted["critical"].(bool)
		suite.True(ok, "failed to convert critical field to bool")

		resultsMap[nodeName] = graph.Container{
			Name:        nodeName,
			Compromised: shared.CompromiseType(compromised),
			Critical:    critical,
		}
	}
	suite.Equal(expectedContainers, resultsMap)
}

// func (suite *VertexTestSuite) TestVertexIdentity() {
// 	g := gremlingo.Traversal_().WithRemote(suite.client)
// 	results, err := g.V().HasLabel(vertex.IdentityLabel).ElementMap().ToList()
// 	suite.NoError(err)
// 	suite.T().Errorf("results: %s", results)
// 	for _, res := range results {
// 		suite.T().Errorf("res: %s", res.String())
// 	}
// }

func (suite *VertexTestSuite) TestVertexNode() {
	g := gremlingo.Traversal_().WithRemote(suite.client)
	results, err := g.V().HasLabel(vertex.NodeLabel).ElementMap().ToList()
	suite.NoError(err)

	suite.Equal(len(expectedNodes), len(results))
	resultsMap := map[string]graph.Node{}
	for _, res := range results {
		res := res.GetInterface()
		converted := res.(map[any]any)

		nodeName, ok := converted["name"].(string)
		suite.True(ok, "failed to convert node name to string")

		compromised, ok := converted["compromised"].(int)
		suite.True(ok, "failed to convert compromised field to CompromiseType")

		isNamespaced, ok := converted["isNamespaced"].(bool)
		suite.True(ok, "failed to convert isNamespaced field to bool")

		namespace, ok := converted["namespace"].(string)
		suite.True(ok, "failed to convert namespace field to string")

		critical, ok := converted["critical"].(bool)
		suite.True(ok, "failed to convert critical field to bool")

		resultsMap[nodeName] = graph.Node{
			Name:         nodeName,
			Compromised:  shared.CompromiseType(compromised),
			IsNamespaced: isNamespaced,
			Namespace:    namespace,
			Critical:     critical,
		}
	}
	suite.Equal(expectedNodes, resultsMap)
}

// func (suite *VertexTestSuite) TestVertexPod() {
// 	g := gremlingo.Traversal_().WithRemote(suite.client)
// 	results, err := g.V().HasLabel(vertex.PodLabel).ElementMap().ToList()
// 	suite.NoError(err)
// 	suite.T().Errorf("results: %s", results)
// 	for _, res := range results {
// 		suite.T().Errorf("res: %s", res.String())
// 	}
// }

// func (suite *VertexTestSuite) TestVertexRole() {
// 	g := gremlingo.Traversal_().WithRemote(suite.client)
// 	results, err := g.V().HasLabel(vertex.RoleLabel).ElementMap().ToList()
// 	suite.NoError(err)
// 	suite.T().Errorf("results: %s", results)
// 	for _, res := range results {
// 		suite.T().Errorf("res: %s", res.String())
// 	}
// }

// func (suite *VertexTestSuite) TestVertexVolume() {
// 	g := gremlingo.Traversal_().WithRemote(suite.client)
// 	results, err := g.V().HasLabel(vertex.VolumeLabel).ElementMap().ToList()
// 	suite.NoError(err)
// 	suite.T().Errorf("results: %s", results)
// 	for _, res := range results {
// 		suite.T().Errorf("res: %s", res.String())
// 	}
// }

func TestVertexTestSuite(t *testing.T) {
	suite.Run(t, new(VertexTestSuite))
}

func (suite *VertexTestSuite) TearDownTest() {
	suite.gdb.Close(context.Background())
}
