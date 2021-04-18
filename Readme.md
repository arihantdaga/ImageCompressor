## ImageCompressor
It is a simple image compressor. I use it for taking backup of my mobile's photos. It uses ImageMagick and reduces file size anywhere from 70-90%, without any visible quality difference. 
It does 2 things - Limits the output file size to 2000x2000 (Width and Height, keeping the ratio same) and set quality to 80%. 

### How to run
Install Imagemagick first
```
https://imagemagick.org/script/download.php
```
If you are using ubuntu like me then follow this - https://www.tecmint.com/install-imagemagick-on-debian-ubuntu/

Update .env
```
INPUT_PATH=/mnt/d/Temp
OUTPUT_PATH=/mnt/d/Temp (You can give OUTPUT_PATH Same as INPUT_PATH also, in that case it will overrite)
```
### Todo
- Support for taking automatic backup also. Incremental, only add new files. 
- Support more configuration