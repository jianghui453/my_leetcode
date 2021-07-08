build:
	g++ -lgtest -pthread -omain_test -std=c++2a main.cc bit/*.cc math_algorithm/*.cc \
	hash/*.cc graph/*.cc bfs/*.cc /usr/local/lib/libgtest.a