Candidate.enterState() {
    self.timer.randomReset()
	self.voteFor = nil
	if self.stopRequests != nil {
		close(self.stopRequests)
	}
	self.stopRequests = nil
}

Candidate.timeout() {
    self.enterState()

    self.currentTerm += 1
    self.voteFor = self.id

    self.stopRequests = make(chan bool)
    go func(){
		voteCount := 0
		for {
			select {
			case <-self.stopRequests:
				return
			default:
				for _, s := range self.server {
					resp := s.sendRequestVote(RequestVote{
						term = self.currentTerm
						candidateID = self.id
						lastLogIndex = len(self.log) - 1
						lastLogTerm = self.log[len(self.log) - 1].term
					})
				    // change to follower
					if resp.term > self.currentTerm {
						state = Follower
						state.enterState()
						return
					}
					if resp.voteGranted {
						voteCount += 1
					}
					// change to leader
					if voteCount > len(self.server) / 2 {
						state = Leader
						state.enterState()
						return
					}
				}
			}
		}
    }()
}