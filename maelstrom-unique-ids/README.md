## Format
The id format is as follows:
```
<machineID>-<unixTimestampMillie>-<AutoIncrementID>
```
### MachineID
This is a random string which is being assigned to the process at the start.
The length is set to 6 with 93 possible characters which results to more than 600 Billion unique machine ids. 

### unixTimestampMillie
This is the request timestamp in Milliseconds. The timestamp in unique ids are generally used to maintains some sort of ordering and uniqueness.

### AutoIncrementID
This is an auto-incrementing integer that resets every 1,000,000. The Unix timestamp is used to ensure that there is at most one reset per millisecond. This means it can generate approximately 1 million IDs per millisecond. Higher loads will result in a wait time for the next millisecond.