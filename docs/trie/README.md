# Trie-Based Route Resolution

This repository presents an implementation of route resolution using the trie data structure. The implementation is designed to resolve HTTP routes, including routes with dynamic path parameters.

## Introduction to Trie

A trie, also known as a prefix tree, is a tree-like data structure that represents a dynamic set of strings, usually associated with an array. Tries are particularly effective for associative array implementations where the keys are strings, as they can exploit common prefixes for efficient operations.

The primary advantage of a trie is its ability to quickly search, insert, and delete strings with time complexity related to the length of the word, rather than the number of words. Additionally, tries offer potential space savings when storing large dictionaries with many common prefixes.

## Route Resolution Implementation

In the context of an HTTP server, routes are usually defined as patterns or paths such as `/users/`, `/users/:id`, etc. Using a trie to store these routes ensures quick resolution of the appropriate handler for a given URL.

The main features of our implementation are:

-   Fast insertion and lookup of routes.
-   Support for dynamic path parameters (e.g., `:id`, `:groupId`).
-   Guarantees to avoid ambiguous routes (e.g., `/users/:id` vs. `/users/:userId`).

Please refer to the code in this repository for detailed implementation.

```
root
│
└─ users
   │
   └─ :id (parameter node for user id)
       │
       └─ groups
           │
           ├─ :groupId (parameter node for group id)
           │   │
           │   └─ approve
           │
           └─ (other static routes if they exist)
       │
       └─ (other static routes if they exist)
   │
   └─ (other static routes if they exist)
│
└─ (other top-level static routes if they exist)

```

## Other Use Cases for Trie

Beyond route resolution, tries are versatile and find use in various domains:

1. **Autocomplete**: Real-time word suggestions as users type.
2. **Spell Checker**: Checking word validity and suggesting corrections.
3. **IP Routing**: Storing IP address routing tables in routers.
4. **Word Games**: Valid word check and finding all possible valid words.
5. **Efficient Storage**: Compact storage of a large dictionary.
6. **Text Analytics**: Tokenizing text, extracting phrases, etc.
7. **DNA Sequence Analysis**: Storing and searching large DNA sequences.

... and many more.

## Conclusion

Tries offer an elegant solution to a range of problems associated with string manipulation and storage. This repository's route resolution implementation showcases the efficiency and clarity that tries bring to the table.

---

Credits to ChatGPT for the help with this README.md 😛
