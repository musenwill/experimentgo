package vellus

/*
Sometimes we need a lot of goroutines to do a job, but don't know how may routines we need
exactly, which often depends on the job size. For example, to calcutate the total size of all files
under a given path recursely, we may alloc goroutines for each directory. Usually we don't know how
many routines we have allocated, so we can hardly wait them to finish synchronizely. Although a
routine pool can help, but that is usually be used in some long run service, it's seemly heavy, and
we just want do something quickly in a short time here.

So here vellus comes, like vellus of Sun Wukong, can become a mass of little monkeys to do dirty
works for you, recycled back to Sun Wukong afterwards.
*/
