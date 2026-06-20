func TestWatchCompactionRace(t *testing.T) {
	// Setup store, create key, delete key, trigger compaction concurrently
	// Assert watcher receives ErrCompacted or the delete event, never skips it.
}