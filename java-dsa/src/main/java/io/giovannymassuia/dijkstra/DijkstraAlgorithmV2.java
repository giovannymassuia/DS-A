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

/*

- https://leetcode.com/problems/network-delay-time/solutions/4191136/bellman/
- https://github.com/wilsonneto-dev/Playing_with_Graphs

// Dijsktra
class Solution:
    def networkDelayTime(self, times: List[List[int]], n: int, k: int) -> int:
        k, _max, edges = k -1, -1, defaultdict(list)
        for s, t, w in times:
            edges[s-1].append((t-1, w))


        costs, priorities = {}, [(0, k)]
        while priorities:
            currCost, curr = heapq.heappop(priorities)
            if curr in costs:
                continue
            costs[curr] = currCost
            _max = max(_max, currCost)

            for target, costToTarget in edges[curr]:
                if target not in costs:
                    heapq.heappush(priorities, (currCost + costToTarget, target))

        if len(costs) != n:
            return -1

        return _max

// Bellman-Ford
    def networkDelayTime(self, times: List[List[int]], n: int, k: int) -> int:
        k = k -1
        costs = [float('inf')] * n
        costs[k] = 0

        for _ in range(n-1):
            change = False
            for source, target, weight in times:
                cost = costs[source-1] + weight
                if costs[target-1] > cost:
                    change = True
                    costs[target-1] = cost
            if not change:
                break

        latency = max(costs)
        if latency == float('inf'):
            return -1
        return latency


// Floyd-Warshall
    def networkDelayTime(self, times: List[List[int]], n: int, k: int) -> int:
        mat = [[float('inf')] * n for _ in range(n)]

        # fill diagnal with zeros
        for i in range(n):
            mat[i][i] = 0

        # fill the edges weights
        for source, target, weight in times:
            mat[source-1][target-1] = weight

        # get k as intermediate candidate
        for intermediate in range(n):
            for i in range(n):
                for j in range(n):
                    if i != j:
                        cost = mat[i][intermediate] + mat[intermediate][j]
                        if mat[i][j] > cost:
                            mat[i][j] = cost

        signalLatency = max(mat[k-1])
        if signalLatency == float('inf'):
            return -1

        return signalLatency

 * */