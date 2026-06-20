func (s *watchableStore) syncWatchers() {
	s.mu.Lock()
	defer s.mu.Unlock()

	curRev := s.kv.Rev()
	compactRev := s.kv.CompactRev()

	for _, w := range s.unsynced {
		if w.minRev <= compactRev {
			w.send(WatchResponse{CompactRevision: compactRev, Err: ErrCompacted})
			w.close()
			continue
		}
		// ... existing logic to retrieve events ...
		// Ensure we check compactRev again after establishing the read transaction
		// to guarantee atomicity.
	}
}