#include <netdb.h>
#include <netinet/in.h>
#include <stddef.h>
#include <stdio.h>
#include <sys/socket.h>
#include <sys/types.h>

int main(int args, char *aargs[]) {
  if (args != 2) {
    fprintf(stderr, "usage: ./showip <host>");
    return 1;
  }

  char *host = aargs[1];
  struct addrinfo hints, *servinfo;

  char ipstr[INET6_ADDRSTRLEN];
}
