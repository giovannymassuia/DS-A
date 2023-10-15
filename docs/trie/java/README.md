## Java Implementation

The trie-based route resolution is implemented in Java, providing an efficient way to manage and resolve HTTP routes, including those with dynamic path parameters.

### Features:

1. **Dynamic Path Parameters**: The trie can handle routes with dynamic path segments, such as `:id` and `:groupId`, capturing these values during the search process for further use.
2. **Ambiguity Handling**: The implementation ensures that ambiguous routes are not allowed. For instance, routes like `/users/:id` and `/users/:userId` can't co-exist, preventing potential confusion during route resolution.

3. **Efficient Insert and Search**: Routes can be inserted and searched with a time and space complexity proportional to the length of the path, ensuring fast operations even with a large number of routes.

### Code Structure:

-   **TrieNode**: This class represents each node in the trie. Each node can have multiple static children (represented as a HashMap) and one parameter child. Additionally, nodes store the route handler and, if they are a parameter node, the parameter's name.

-   **Trie**: The primary class that offers `insert` and `search` methods. The insert method allows adding routes to the trie, while the search method resolves a given path to its corresponding handler, capturing any path parameters in the process.

To use the implementation, simply create an instance of the Trie class, add routes using the insert method, and resolve routes with the search method.

For detailed method functionalities and complexities, refer to the sections above.

## Time and Space Complexity Analysis

### Insert Method

**Time Complexity:**

-   Splitting the path string into segments based on the delimiter `/`: \(O(n)\), where \(n\) is the length of the path.
-   Traversing and inserting each segment into the trie: \(O(m)\), where \(m\) is the number of segments. Note that this is bounded by \(n\) because in the worst case every character might be a segment (though this would be unusual).
-   Thus, the overall time complexity for the insert operation is \(O(n)\).

**Space Complexity:**

-   The string array to hold segments: \(O(m)\) which is \(O(n)\) in the worst case.
-   The trie itself grows based on the number of unique paths inserted. Each path can take up to \(O(n)\) space (because of nodes for each character/segment and possible pointers).
-   For a single insert operation, excluding the size of the trie, the space complexity is \(O(n)\).

### Search Method

**Time Complexity:**

-   Splitting the path string into segments: \(O(n)\), where \(n\) is the length of the path.
-   Traversing the trie to find a match for each segment: \(O(m)\), where \(m\) is the number of segments. Again, \(m\) is bounded by \(n\).
-   Therefore, the overall time complexity for the search operation is also \(O(n)\).

**Space Complexity:**

-   The string array to hold segments: \(O(m)\) which is \(O(n)\) in the worst case.
-   A map (or dictionary) to hold captured parameters: In the worst case, all segments are parameters, so \(O(m)\) which again is \(O(n)\) in the extreme.
-   Thus, for a single search operation, the space complexity is \(O(n)\).

In conclusion, both the `insert` and `search` methods in our trie implementation for route resolution have a time and space complexity of \(O(n)\), where \(n\) is the length of the path being inserted or searched for.

---

Credits to ChatGPT for the help with this README.md 😛
