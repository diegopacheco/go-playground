### Run
```bash
./run.sh
```

### Test
```bash
 go run main.go "SELECT * FROM table WHERE attr = 10"
```
Output:
```
&main.Select{
  Expression: &main.SelectExpression{
    All: true,
  },
  From: &main.From{
    TableExpressions: []*main.TableExpression{
      {
        Table: "table",
      },
    },
    Where: &main.Expression{
      Or: []*main.OrCondition{
        {
          And: []*main.Condition{
            {
              Operand: &main.ConditionOperand{
                Operand: &main.Operand{
                  Summand: []*main.Summand{
                    {
                      LHS: &main.Factor{
                        LHS: &main.Term{
                          SymbolRef: &main.SymbolRef{
                            Symbol: "attr",
                          },
                        },
                      },
                    },
                  },
                },
                ConditionRHS: &main.ConditionRHS{
                  Compare: &main.Compare{
                    Operator: "=",
                    Operand: &main.Operand{
                      Summand: []*main.Summand{
                        {
                          LHS: &main.Factor{
                            LHS: &main.Term{
                              Value: &main.Value{
                                Number: &10,
                              },
                            },
                          },
                        },
                      },
                    },
                  },
                },
              },
            },
          },
        },
      },
    },
  },
}
```
