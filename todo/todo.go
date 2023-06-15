package todo


/*
Named types
  - can be any know type (struct, string, int, slice, a new type we declared, etc)
  - methods can be declared on it
  - not an alias - Explicit type
*/
type Item struct {
	/* EXPORTED: */
	Text string
}