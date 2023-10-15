package io.giovannymassuia.dijkstra;

import io.giovannymassuia.dijkstra.models.Distance;
import io.giovannymassuia.dijkstra.models.Edge;
import io.giovannymassuia.dijkstra.models.Node;
import io.giovannymassuia.dijkstra.models.NodeKey;

import java.util.*;

public class DijkstraAlgorithmV1 implements DijkstraAlgorithm {

    @Override
    public Map<NodeKey, Distance> calculateDistances(Map<NodeKey, Node> graph, NodeKey start, NodeKey end) {
        var visited = new HashSet<NodeKey>();
        var distances = new HashMap<NodeKey, Distance>();
        var queue = new PriorityQueue<>(Comparator.comparingInt(Distance::value));

        queue.add(Distance.of(0, start, null));

        while (visited.size() < graph.size() && !queue.isEmpty()) {
            var current = queue.poll();

            if (visited.contains(current.key())) continue;

            // early exit
            if (current.key().equals(end)) return distances;

            visited.add(current.key());

            for (Edge edge : graph.get(current.key()).edges()) {
                if (visited.contains(edge.key())) {
                    continue;
                }

                var currDistance = distances.getOrDefault(edge.key(), Distance.of(Integer.MAX_VALUE, edge.key(), null));
                var newDistance = Distance.of(current.value() + edge.distance(), edge.key(), current.key());

                if (newDistance.value() < currDistance.value()) {
                    distances.put(edge.key(), newDistance);
                    queue.add(newDistance);
                }
            }
        }

        return distances;
    }

}
