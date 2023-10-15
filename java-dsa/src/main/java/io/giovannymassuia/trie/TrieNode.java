package io.giovannymassuia.trie;

import java.util.HashMap;
import java.util.Map;
import java.util.function.Supplier;

public class TrieNode {
    public Map<String, TrieNode> staticChildren = new HashMap<>();
    public TrieNode parameterChild = null;
    public Supplier<Boolean> handler; // This could be some kind of route handler.
    public String parameterName = null;  // If null, it's a static node. Otherwise, it's a parameterized node.
}