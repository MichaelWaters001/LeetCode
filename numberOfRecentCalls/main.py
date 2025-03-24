class RecentCounter:

    def __init__(self):
        self.calls = []
        self.prt = 0

    def ping(self, t: int) -> int:
        self.calls.append(t)

        while self.calls[self.ptr] < t - 3000:
            self.ptr += 1

        return len(self.calls) - self.ptr


# Your RecentCounter object will be instantiated and called as such:
# obj = RecentCounter()
# param_1 = obj.ping(t)