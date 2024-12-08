{:ok, body} = File.read("./test.txt")

x = String.split(body)

IO.puts(x)
