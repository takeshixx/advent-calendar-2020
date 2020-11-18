import sys
import socket
import ssl
from socket import socket, AF_INET, SOCK_DGRAM, SOL_SOCKET, SO_REUSEADDR
from dtls import do_patch
do_patch()

from threading import Thread
from select import select

BIND_WEBSERVER = ('0.0.0.0', 57000)
BUFSIZE = 4096

def handle_dtls_client(client_sock, target, cfg):
	server_sock = ssl.wrap_socket(socket(AF_INET, SOCK_DGRAM), cfg.cert.name, cfg.key.name)
	server_sock.connect(target)
	do_relay_dtls(client_sock, server_sock, cfg)


def create_server(relay, cfg):
    serv = ssl.wrap_socket(socket(AF_INET, SOCK_DGRAM), cfg.clientcert.name, cfg.clientkey.name)
    serv.setsockopt(SOL_SOCKET, SO_REUSEADDR, 1)
    serv.bind((cfg.listen, lport))
    serv.listen(2)

    print '[+] Relay listening on %s %d -> %s:%d' % relay
    while True:
        if proto == 'udp':
            client, addr = serv.accept()
            dest_str = '%s:%d' % (relay[2], relay[3])

            print '[+] New client:', addr, "->", color(dest_str, 4)
            thread = Thread(target=handle_dtls_client, args=(client, (rhost, rport), cfg))
            thread.start()

def main():
	parser = argparse.ArgumentParser(
		description='%s version %.2f' % (__prog_name__, __version__))

	parser.add_argument('-l', '--listen',
		action='store',
		metavar='<listen>',
		dest='listen',
		help='Address the relays will listen on. Default: 0.0.0.0',
		default='0.0.0.0')

	parser.add_argument('-c', '--cert',
		action='store',
		metavar='<cert>',
		dest='cert',
		type=argparse.FileType('r'),
		help='Certificate file to use for SSL/TLS interception',
		default=False)

	parser.add_argument('-k', '--key',
		action='store',
		metavar='<key>',
		dest='key',
		type=argparse.FileType('r'),
		help='Private key file to use for SSL/TLS interception',
		default=False)

	cfg = parser.parse_args()
	cfg.prog_name = __prog_name__

	relays = [item for sublist in cfg.relays for item in sublist]

	cfg.relays = []
	for r in relays:
		r = r.split(':')

		try:
			if len(r) == 3:
				cfg.relays.append(('tcp', int(r[0]), r[1], int(r[2])))
			elif len(r) == 4 and r[0] in ['tcp', 'udp']:
				cfg.relays.append((r[0], int(r[1]), r[2], int(r[3])))
			else:
				raise

			if r[0] == 'udp' and cfg.listen.startswith('127.0.0'):
				print color("[!] In UDP, it's not recommended to bind to 127.0.0.1. If you see errors, try to bind to your LAN IP address instead.", 1)

		except:
			sys.exit('[!] error: Invalid relay specification, see help.')

	if not (cfg.cert and cfg.key):
		print color("[!] Server cert/key not provided, SSL/TLS interception will not be available.", 1)

	if not (cfg.clientcert and cfg.clientkey):
		print color("[!] Client cert/key not provided.", 1)

	# There is no point starting the local web server
	# if we are not going to intercept the req/resp (monitor only).
	if cfg.proxy:
		start_ws()
	else:
		print color("[!] Interception disabled! %s will run in monitoring mode only." % __prog_name__, 1)

	# If a script was specified, import it
	if cfg.script:
		try:
			from imp import load_source
			cfg.script_module = load_source(cfg.script.name, cfg.script.name)

		except Exception as e:
			print color("[!] %s" % str(e))
			sys.exit()
	# If a ssl keylog file was specified, dump (pre-)master secrets
	if cfg.sslkeylog:
		try:
			import sslkeylog
			sslkeylog.set_keylog(cfg.sslkeylog)

		except Exception as e:
			print color("[!] %s" % str(e))
			sys.exit()

	server_threads = []
	for relay in cfg.relays:
		server_threads.append(Thread(target=create_server, args=(relay, cfg)))

	for t in server_threads:
		t.setDaemon(True)
		t.start()
		time.sleep(.2)

	while True:
		try:
			time.sleep(100)

		except KeyboardInterrupt:
			sys.exit("\rExiting...")

if __name__=='__main__':
	main()