# typewalk
List and count all file types (recursively) in a directory

## Usage:

### recursively walk
```> ./typewalk.exe walk --path "C:\\"```

### merge two files
```> ./typewalk.exe merge --file1 "host1.json" --file2 "host2.json"```

## Example json output
```
{
  "folder": "Windows\\BitLockerDiscoveryVolumeContents",
  "file_count": 41,
  "file_types": {
   ".exe": 1,
   ".inf": 1,
   ".mui": 38,
   ".url": 1
  },
  "additional_info": ""
 },
 {
  "folder": "Windows\\Boot",
  "file_count": 204,
  "file_types": {
   "": 4,
   ".bin": 3,
   ".com": 1,
   ".dll": 17,
   ".efi": 3,
   ".exe": 1,
   ".ini": 1,
   ".mui": 154,
   ".p7b": 1,
   ".sdi": 2,
   ".stl": 1,
   ".ttf": 16
  },
  "additional_info": ""
 },
