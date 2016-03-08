setup() {
  export MAGIC_SEAWEED_API_KEY=abc123
  export MAGIC_SEAWEED_API_SECRET=abc123
  export MAGIC_SEAWEED_CACHE_DIR=test/tmp
  ew="bin/seaweed-cli"
  mkdir -p test/tmp
}

teardown() {
  rm test/tmp/*
  echo $output
}

fixture() {
  cp test/{fixtures,tmp}/seaweed_391
  today=`date +%s`
  sed -i ".bak" "s/today_timestamp/$today/g" test/tmp/seaweed_391
}
