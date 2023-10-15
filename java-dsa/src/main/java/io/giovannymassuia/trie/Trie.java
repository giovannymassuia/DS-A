package io.giovannymassuia.trie;

import java.util.HashMap;
import java.util.HashSet;
import java.util.Map;
import java.util.Set;
import java.util.function.Supplier;

public class Trie {
    private final TrieNode root = new TrieNode();

    public void insert(String path, Supplier<Boolean> handler) {
        TrieNode current = root;
        String[] segments = path.split("/");
        Set<String> seenParameters = new HashSet<>();

        for (String segment : segments) {
            if (segment.startsWith(":")) {
                String paramName = segment.substring(1);

                if (seenParameters.contains(paramName)) {
                    throw new IllegalArgumentException("Cannot use the same parameter name more than once in a route.");
                }
                seenParameters.add(paramName);

                if (current.parameterChild != null) {
                    // Check if the parameter is of the same type.
                    if (!current.parameterChild.parameterName.equals(paramName)) {
                        throw new IllegalArgumentException("Invalid path: " + path + ". Parameter " + paramName + " is already defined.");
                    }
                } else {
                    TrieNode paramNode = new TrieNode();
                    paramNode.parameterName = paramName;
                    current.parameterChild = paramNode;
                }
                current = current.parameterChild;
            } else {
                current.staticChildren.putIfAbsent(segment, new TrieNode());
                current = current.staticChildren.get(segment);
            }
        }
        current.handler = handler;
    }

    public SearchResult search(String path) {
        TrieNode current = root;
        String[] segments = path.split("/");
        Map<String, String> capturedParams = new HashMap<>();

        for (String segment : segments) {
            TrieNode next = current.staticChildren.get(segment);

            // If exact match is not found, check if a parameter child exists.
            if (next == null && current.parameterChild != null) {
                next = current.parameterChild;
                capturedParams.put(next.parameterName, segment);
            }

            if (next == null) {
                return new SearchResult(() -> false, null);
            }

            current = next;
        }
        return new SearchResult(current.handler, capturedParams);
    }

    public record SearchResult(Supplier<Boolean> handler, Map<String, String> params) {
    }

}
