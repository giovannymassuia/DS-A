package io.giovannymassuia.dijkstra;

import io.giovannymassuia.dijkstra.models.Edge;
import io.giovannymassuia.dijkstra.models.Node;
import io.giovannymassuia.dijkstra.models.NodeKey;

import java.util.HashMap;
import java.util.Map;

public class GraphsTest {

    static Map<NodeKey, Node> graph1() {
        /* {
            ['A'] = new() { ['B'] = 5, ['C'] = 1 },
            ['B'] = new() { ['D'] = 2, ['B'] = 4, ['E'] = 4 },
            ['C'] = new() { ['E'] = 5, ['B'] = 2, ['D'] = 5 },
            ['D'] = new() { ['E'] = 2 },
            ['E'] = new()
        }; */

        return Map.of(
                NodeKey.of("A"),
                Node.of("A").addEdges(
                        Edge.of(NodeKey.of("B"), 5),
                        Edge.of(NodeKey.of("C"), 1)),

                NodeKey.of("B"),
                Node.of("B").addEdges(
                        Edge.of(NodeKey.of("D"), 2),
                        Edge.of(NodeKey.of("C"), 4),
                        Edge.of(NodeKey.of("E"), 4)),

                NodeKey.of("C"),
                Node.of("C").addEdges(Edge.of(NodeKey.of("E"), 5),
                        Edge.of(NodeKey.of("B"), 2),
                        Edge.of(NodeKey.of("D"), 5)),

                NodeKey.of("D"),
                Node.of("D").addEdges(
                        Edge.of(NodeKey.of("E"), 2)),

                NodeKey.of("E"),
                Node.of("E")
        );
    }

    static Map<NodeKey, Node> graph2() {
        // {
        //     ['A'] = new() { ['B'] = 3, ['D'] = 4, ['S'] = 7 },
        //     ['B'] = new() { ['A'] = 3, ['D'] = 4, ['S'] = 2, ['H'] = 1 },
        //     ['C'] = new() { ['S'] = 3, ['L'] = 2 },
        //     ['D'] = new() { ['A'] = 4, ['B'] = 4, ['F'] = 5 },
        //     ['E'] = new() { ['G'] = 2, ['K'] = 5 },
        //     ['F'] = new() { ['D'] = 5, ['H'] = 3 },
        //     ['G'] = new() { ['E'] = 2, ['H'] = 2 },
        //     ['H'] = new() { ['B'] = 1, ['F'] = 3, ['G'] = 2 },
        //     ['I'] = new() { ['J'] = 6, ['K'] = 4, ['L'] = 4 },
        //     ['J'] = new() { ['I'] = 6, ['K'] = 4, ['L'] = 4 },
        //     ['K'] = new() { ['E'] = 5 ,['I'] = 4, ['J'] = 4 },
        //     ['L'] = new() { ['C'] = 2, ['I'] = 4, ['J'] = 4 },
        //     ['S'] = new() { ['A'] = 7, ['B'] = 2, ['C'] = 3 }
        // };

        var graph = new HashMap<NodeKey, Node>();

        graph.put(NodeKey.of("A"), Node.of("A").addEdges(
                Edge.of(NodeKey.of("B"), 3),
                Edge.of(NodeKey.of("D"), 4),
                Edge.of(NodeKey.of("S"), 7)));

        graph.put(NodeKey.of("B"), Node.of("B").addEdges(
                Edge.of(NodeKey.of("A"), 3),
                Edge.of(NodeKey.of("D"), 4),
                Edge.of(NodeKey.of("S"), 2),
                Edge.of(NodeKey.of("H"), 1)));

        graph.put(NodeKey.of("C"), Node.of("C").addEdges(
                Edge.of(NodeKey.of("S"), 3),
                Edge.of(NodeKey.of("L"), 2)));

        graph.put(NodeKey.of("D"), Node.of("D").addEdges(
                Edge.of(NodeKey.of("A"), 4),
                Edge.of(NodeKey.of("B"), 4),
                Edge.of(NodeKey.of("F"), 5)));

        graph.put(NodeKey.of("E"), Node.of("E").addEdges(
                Edge.of(NodeKey.of("G"), 2),
                Edge.of(NodeKey.of("K"), 5)));

        graph.put(NodeKey.of("F"), Node.of("F").addEdges(
                Edge.of(NodeKey.of("D"), 5),
                Edge.of(NodeKey.of("H"), 3)));

        graph.put(NodeKey.of("G"), Node.of("G").addEdges(
                Edge.of(NodeKey.of("E"), 2),
                Edge.of(NodeKey.of("H"), 2)));

        graph.put(NodeKey.of("H"), Node.of("H").addEdges(
                Edge.of(NodeKey.of("B"), 1),
                Edge.of(NodeKey.of("F"), 3),
                Edge.of(NodeKey.of("G"), 2)));

        graph.put(NodeKey.of("I"), Node.of("I").addEdges(
                Edge.of(NodeKey.of("J"), 6),
                Edge.of(NodeKey.of("K"), 4),
                Edge.of(NodeKey.of("L"), 4)));

        graph.put(NodeKey.of("J"), Node.of("J").addEdges(
                Edge.of(NodeKey.of("I"), 6),
                Edge.of(NodeKey.of("K"), 4),
                Edge.of(NodeKey.of("L"), 4)));

        graph.put(NodeKey.of("K"), Node.of("K").addEdges(
                Edge.of(NodeKey.of("E"), 5),
                Edge.of(NodeKey.of("I"), 4),
                Edge.of(NodeKey.of("J"), 4)));

        graph.put(NodeKey.of("L"), Node.of("L").addEdges(
                Edge.of(NodeKey.of("C"), 2),
                Edge.of(NodeKey.of("I"), 4),
                Edge.of(NodeKey.of("J"), 4)));

        graph.put(NodeKey.of("S"), Node.of("S").addEdges(
                Edge.of(NodeKey.of("A"), 7),
                Edge.of(NodeKey.of("B"), 2),
                Edge.of(NodeKey.of("C"), 3)));

        return graph;
    }

}
