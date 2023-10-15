package io.giovannymassuia.dijkstra;

import io.giovannymassuia.dijkstra.models.Distance;
import io.giovannymassuia.dijkstra.models.Node;
import io.giovannymassuia.dijkstra.models.NodeKey;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

import java.util.List;
import java.util.Map;

import static io.giovannymassuia.dijkstra.GraphsTest.graph1;
import static io.giovannymassuia.dijkstra.GraphsTest.graph2;
import static org.junit.jupiter.api.Assertions.assertEquals;

class DijkstraAlgorithmTest {

    private DijkstraAlgorithm dijkstraAlgorithmV1 = new DijkstraAlgorithmV1();
    private DijkstraAlgorithm dijkstraAlgorithmV2 = new DijkstraAlgorithmV2();

    @BeforeEach
    public void setUp() {
        dijkstraAlgorithmV1 = new DijkstraAlgorithmV1();
        dijkstraAlgorithmV2 = new DijkstraAlgorithmV2();
    }

    @Test
    public void testShortestPath_1() {
        var graph = graph1();

        List<Node> shortestPathV1 = dijkstraAlgorithmV1.shortestPath(graph, NodeKey.of("A"), NodeKey.of("D"));
        List<Node> shortestPathV2 = dijkstraAlgorithmV2.shortestPath(graph, NodeKey.of("A"), NodeKey.of("D"));

        List.of(shortestPathV1, shortestPathV2).forEach(shortestPath -> {
            assertEquals(4, shortestPath.size());
            assertEquals("A", shortestPath.get(0).key().value());
            assertEquals("C", shortestPath.get(1).key().value());
            assertEquals("B", shortestPath.get(2).key().value());
            assertEquals("D", shortestPath.get(3).key().value());
        });
    }

    @Test
    public void testShortestPath_2() {
        var graph = graph2();

        List<Node> shortestPathV1 = dijkstraAlgorithmV1.shortestPath(graph, NodeKey.of("S"), NodeKey.of("E"));
        List<Node> shortestPathV2 = dijkstraAlgorithmV2.shortestPath(graph, NodeKey.of("S"), NodeKey.of("E"));

        List.of(shortestPathV1, shortestPathV2).forEach(shortestPath -> {
            assertEquals(5, shortestPath.size());
            assertEquals("S", shortestPath.get(0).key().value());
            assertEquals("B", shortestPath.get(1).key().value());
            assertEquals("H", shortestPath.get(2).key().value());
            assertEquals("G", shortestPath.get(3).key().value());
            assertEquals("E", shortestPath.get(4).key().value());
        });

        shortestPathV1 = dijkstraAlgorithmV1.shortestPath(graph, NodeKey.of("A"), NodeKey.of("I"));
        shortestPathV2 = dijkstraAlgorithmV2.shortestPath(graph, NodeKey.of("A"), NodeKey.of("I"));

        List.of(shortestPathV1, shortestPathV2).forEach(shortestPath -> {
            assertEquals(6, shortestPath.size());
            assertEquals("A", shortestPath.get(0).key().value());
            assertEquals("B", shortestPath.get(1).key().value());
            assertEquals("S", shortestPath.get(2).key().value());
            assertEquals("C", shortestPath.get(3).key().value());
            assertEquals("L", shortestPath.get(4).key().value());
            assertEquals("I", shortestPath.get(5).key().value());
        });
    }

    @Test
    public void testShortestDistance_1() {
        var graph = graph1();

        List<Distance> testDistances = List.of(
                Distance.of(5, NodeKey.of("A"), NodeKey.of("D")),
                Distance.of(6, NodeKey.of("A"), NodeKey.of("E")),
                Distance.of(4, NodeKey.of("B"), NodeKey.of("C")),
                Distance.of(-1, NodeKey.of("E"), NodeKey.of("A"))
        );

        assertDistances(graph, testDistances);
    }

    @Test
    public void testShortestDistance_2() {
        var graph = graph2();

        List<Distance> testDistances = List.of(
                Distance.of(7, NodeKey.of("S"), NodeKey.of("E")),
                Distance.of(5, NodeKey.of("S"), NodeKey.of("L")),
                Distance.of(14, NodeKey.of("A"), NodeKey.of("I"))
        );

        assertDistances(graph, testDistances);
    }

    private void assertDistances(Map<NodeKey, Node> graph, List<Distance> testDistances) {
        testDistances.forEach(testDist -> {
            int v1Result = dijkstraAlgorithmV1.shortestDistance(graph, testDist.key(), testDist.previous());
            int v2Result = dijkstraAlgorithmV2.shortestDistance(graph, testDist.key(), testDist.previous());

            assertEquals(testDist.value(), v1Result);
            assertEquals(testDist.value(), v2Result);
        });
    }
}