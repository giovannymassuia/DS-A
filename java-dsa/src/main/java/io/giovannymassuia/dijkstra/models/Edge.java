package io.giovannymassuia.dijkstra.models;

public record Edge(NodeKey key, int distance) {
    public static Edge of(NodeKey key, int distance) {
        return new Edge(key, distance);
    }
}