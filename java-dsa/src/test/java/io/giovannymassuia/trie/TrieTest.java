package io.giovannymassuia.trie;

import org.junit.jupiter.api.Test;

import java.util.Map;

import static org.junit.jupiter.api.Assertions.*;

class TrieTest {

    @Test
    public void test() {
        Trie trie = new Trie();

        trie.insert("/users", () -> true);
        trie.insert("/users/:userId", () -> true);
        trie.insert("/users/:userId/profile", () -> true);
        trie.insert("/users/:userId/groups/:groupId", () -> true);
        trie.insert("/users/:userId/groups/:groupId/approve", () -> true);
        trie.insert("/auth/login", () -> true);
        trie.insert("/auth/logout", () -> true);

        var result = trie.search("/users");
        assertValidHandlerEmptyParams(result);

        result = trie.search("/users/1");
        assertValidHandlerWithParams(result, Map.of("userId", "1"));

        result = trie.search("/users/1/profile");
        assertValidHandlerWithParams(result, Map.of("userId", "1"));

        result = trie.search("/users/1/groups/2");
        assertValidHandlerWithParams(result, Map.of("userId", "1", "groupId", "2"));

        result = trie.search("/users/1/groups/2/approve");
        assertValidHandlerWithParams(result, Map.of("userId", "1", "groupId", "2"));

        result = trie.search("/auth/login");
        assertValidHandlerEmptyParams(result);

        result = trie.search("/auth/logout");
        assertValidHandlerEmptyParams(result);

        result = trie.search("/invalid-path");
        assertInvalidResult(result);
    }

    @Test
    public void testInvalid() {
        Trie trie = new Trie();

        trie.insert("/users", () -> true);
        trie.insert("/users/:id", () -> true);

        // should not work because there is already a path, just with a different parameter name
        var exception = assertThrows(IllegalArgumentException.class, () -> trie.insert("/users/:userId/approve", () -> true));
        assertEquals("Invalid path: /users/:userId/approve. Parameter userId is already defined.", exception.getMessage());

        // should not work because there is a variable with the same name (id)
        exception = assertThrows(IllegalArgumentException.class, () -> trie.insert("/users/:id/group/:id", () -> true));
        assertEquals("Cannot use the same parameter name more than once in a route.", exception.getMessage());

    }

    private void assertValidHandlerEmptyParams(Trie.SearchResult result) {
        assertNotNull(result.handler());
        assertTrue(result.handler().get());
        assertTrue(result.params().isEmpty());
    }

    private void assertValidHandlerWithParams(Trie.SearchResult result, Map<String, String> expectedParams) {
        assertNotNull(result.handler());
        assertTrue(result.handler().get());

        var params = result.params();
        assertEquals(expectedParams.size(), params.size());
        for (var entry : expectedParams.entrySet()) {
            assertEquals(entry.getValue(), params.get(entry.getKey()));
        }
    }

    private void assertInvalidResult(Trie.SearchResult result) {
        assertNotNull(result.handler());
        assertFalse(result.handler().get());
        assertNull(result.params());
    }

}