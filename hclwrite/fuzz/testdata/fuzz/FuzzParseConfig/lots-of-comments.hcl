go test fuzz v1
[]byte("// comment\nblock {\n  // another comment\n  another_block { # comment\n    // comment\n    foo : bar\n  }\n\n  /* commented out block\n  blah {\n    bar : foo\n  }\n  */\n}\n")
