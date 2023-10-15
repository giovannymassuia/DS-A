package io.giovannymassuia.dijkstra;

import io.giovannymassuia.dijkstra.models.Distance;
import io.giovannymassuia.dijkstra.models.Node;
import io.giovannymassuia.dijkstra.models.NodeKey;

import java.util.Comparator;
import java.util.HashMap;
import java.util.Map;
import java.util.PriorityQueue;

public class DijkstraAlgorithmV2 implements DijkstraAlgorithm {

    @Override
    public Map<NodeKey, Distance> calculateDistances(Map<NodeKey, Node> graph, NodeKey start, NodeKey end) {
        var distances = new HashMap<NodeKey, Distance>();
        var queue = new PriorityQueue<>(Comparator.comparingInt(Distance::value));

        queue.add(Distance.of(0, start, null));

        while (!queue.isEmpty()) {
            var current = queue.poll();

            if (distances.containsKey(current.key())) continue;

            distances.put(current.key(), current);

            // early exit
//            if (current.key().equals(end)) return distances;

            graph.get(current.key()).edges().stream()
                    .filter(edge -> !distances.containsKey(edge.key()))
                    .forEach(edge -> queue.add(Distance.of(current.value() + edge.distance(), edge.key(), current.key())));
        }

        return distances;
    }

}
