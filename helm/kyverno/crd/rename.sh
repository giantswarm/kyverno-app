for crd in $(ls | grep crd)
do
  mv $crd $(echo $crd | awk -F "." '{print $1 ".yaml"}')
done
