### Run

```bash
./run.sh
```

### Result

```bash
❯ ./run.sh
🌳 Tree-Sitter Analysis
=========================================
📝 Source Code: const foo = 1 + 2
🔍 Language: JavaScript
📊 Node Count: 1

🔧 Raw S-Expression:
(program (lexical_declaration (variable_declarator name: (identifier) value: (binary_expression left: (number) right: (number)))))

🌲 Pretty Tree Structure:
└── program 📄 "const foo = 1 + 2" [1:1-1:18]
    └── lexical_declaration 📄 "const foo = 1 + 2" [1:1-1:18]
        └── variable_declarator 📄 "foo = 1 + 2" [1:7-1:18]
            ├── identifier 📄 "foo" [1:7-1:10]
            └── binary_expression 📄 "1 + 2" [1:13-1:18]
                ├── number 📄 "1" [1:13-1:14]
                └── number 📄 "2" [1:17-1:18]

📋 Node Details:
🔸 program
   📍 Position: 1:1 to 1:18
   📏 Bytes: 0-17
   📝 Text: "const foo = 1 + 2"

  🔸 lexical_declaration
     📍 Position: 1:1 to 1:18
     📏 Bytes: 0-17
     📝 Text: "const foo = 1 + 2"

    🔸 variable_declarator
       📍 Position: 1:7 to 1:18
       📏 Bytes: 6-17
       📝 Text: "foo = 1 + 2"

      🔸 identifier
         📍 Position: 1:7 to 1:10
         📏 Bytes: 6-9
         📝 Text: "foo"

      🔸 binary_expression
         📍 Position: 1:13 to 1:18
         📏 Bytes: 12-17
         📝 Text: "1 + 2"

        🔸 number
           📍 Position: 1:13 to 1:14
           📏 Bytes: 12-13
           📝 Text: "1"

        🔸 number
           📍 Position: 1:17 to 1:18
           📏 Bytes: 16-17
           📝 Text: "2"
```