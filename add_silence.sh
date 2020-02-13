mkdir ./slow
for i in *.mp3;
  do name=`echo "$i" | cut -d'.' -f1`
  echo "$name"
  ffmpeg -i "$i" -filter:a "atempo=0.7" "./slow/${name}_.mp3"
  ffmpeg -i "./slow/${name}_.mp3" -af "apad=pad_dur=1" "./slow/${name}_slow.mp3"
  ffmpeg -i "$i" -af "apad=pad_dur=1" "./slow/$i"
  rm "./slow/${name}_.mp3"
done
