#include "ports.h"

#if defined(OS_WIN) // not tested
#include <windows.h>
#else
// linux//UNIX
#include <unistd.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#endif

// find a free port to open socket on
// returns 0 if failed
int findFreePort() {
    #if defined(OS_WIN) // not tested
        WSADATA wsaData;
        if (WSAStartup(MAKEWORD(2, 2), &wsaData) != 0)
            return 0;
    
        SOCKET sock = socket(AF_INET, SOCK_STREAM, 0);
        if (sock == INVALID_SOCKET) {
            WSACleanup();
            return 0;
        }
    #else
        int sock = socket(AF_INET, SOCK_STREAM, 0);
        if (sock < 0) return 0;
    #endif
    
        sockaddr_in addr{};
        addr.sin_family = AF_INET;
        addr.sin_addr.s_addr = htonl(INADDR_ANY);
        addr.sin_port = 0;
    
        if (bind(sock, reinterpret_cast<sockaddr*>(&addr), sizeof(addr)) != 0) {
    #if defined(OS_WIN)
            closesocket(sock);
            WSACleanup();
    #else
            close(sock);
    #endif
            return 0;
        }
    
        socklen_t len = sizeof(addr);
        if (getsockname(sock, reinterpret_cast<sockaddr*>(&addr), &len) != 0) {
    #if defined(OS_WIN)
            closesocket(sock);
            WSACleanup();
    #else
            close(sock);
    #endif
            return 0;
        }
    
        int port = ntohs(addr.sin_port);
    
    #if defined(OS_WIN)
        closesocket(sock);
        WSACleanup();
    #else
        close(sock);
    #endif
    
        return port;
    }    