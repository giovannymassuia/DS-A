package io.giovannymassuia.dijkstra;

import io.giovannymassuia.dijkstra.models.Distance;
import io.giovannymassuia.dijkstra.models.Node;
import io.giovannymassuia.dijkstra.models.NodeKey;

import java.util.*;

public interface DijkstraAlgorithm {

    /**
     * Given a graph, a start node and an end node, returns a map of distances from the start until end node.
     *
     * @param graph the graph
     * @param start the start node
     * @param end   the end node
     * @return a map of distances from the start node to each node
     */
    Map<NodeKey, Distance> calculateDistances(Map<NodeKey, Node> graph, NodeKey start, NodeKey end);

    /**
     * Given a graph, a start node and an end node, returns the shortest path between them.
     *
     * @param graph the graph
     * @param start the start node
     * @param end   the end node
     * @return the shortest path between the start and end nodes
     */
    default List<Node> shortestPath(Map<NodeKey, Node> graph, NodeKey start, NodeKey end) {
        var distances = this.calculateDistances(graph, start, end);

        var shortestPath = new ArrayList<Node>();
        var current = end;
        while (current != null) {
            shortestPath.add(graph.get(current));
            current = Optional.ofNullable(distances.get(current)).map(Distance::previous).orElse(null);
        }
        Collections.reverse(shortestPath);

        return shortestPath;
    }

    /**
     * Given a graph, a start node and an end node, returns the shortest distance between them.
     *
     * @param graph the graph
     * @param start the start node
     * @param end   the end node
     * @return the shortest distance between the start and end nodes
     */
    default int shortestDistance(Map<NodeKey, Node> graph, NodeKey start, NodeKey end) {
        var distances = this.calculateDistances(graph, start, end);
        return Optional.ofNullable(distances.get(end)).map(Distance::value).orElse(-1);

    }

}
