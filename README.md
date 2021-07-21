# HealthyFish

Calculate pack size based on the following rules:
1. Only whole packs can be sent. Packs cannot be broken open.
2. Within the constraints of Rule 1 above, send out no more items than necessary to fulfil the order.
3. Within the constraints of Rules 1 & 2 above, send out as few packs as possible to fulfil each order.

Example Table:
  
| Items ordered        | Correct number of packs           | Incorrect number of packs  |
| ------------- |-------------| -----|
| 1      | 1 x 250 | 1 x 500 – more items than necessary |
| 250      | 1 x 250      |   1 x 500 – more items than necessary |
| 251 | 1 x 500      |    2 x 250 – more packs than necessary |
| 501 | 1 x 500 <br /> 1 x 250    |    1 x 1000 – more items than necessary <br /> 3 x 250 – more packs than necessary |
| 12001 | 2 x 5000 <br /> 1 x 2000 <br /> 1 x 250      |    3 x 5000 – more items than necessary |

## Notes

- Using in-memory array as a database 
- First Go programme and the directory layout could be improved
- Use of AWS CDK deployment caused some issues with github actions for CI/CD
