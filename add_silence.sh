mkdir ./silent
for i in *.mp3;
  do name=`echo "$i" | cut -d'.' -f1`
  echo "$name"
  ffmpeg -i "$i" -filter:a "atempo=0.7" "./silent/${name}_.mp3"
  ffmpeg -i "./silent/${name}_.mp3" -af "apad=pad_dur=1" "./silent/${name}_slow.mp3"
done
