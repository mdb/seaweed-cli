load test_helper

@test "today's surf" {
  fixture response.json

  run $ew today 391

  echo $output

  echo $output | grep "7.50ft"
}
