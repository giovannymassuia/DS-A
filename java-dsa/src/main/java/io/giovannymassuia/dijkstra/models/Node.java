package io.giovannymassuia.dijkstra.models;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

public record Node(NodeKey key, List<Edge> edges) {
    public static Node of(String key) {
        return new Node(NodeKey.of(key), new ArrayList<>());
    }

    public Node addEdges(Edge... newEdges) {
        Collections.addAll(edges, newEdges);
        return this;
    }
}