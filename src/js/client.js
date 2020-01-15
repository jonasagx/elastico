const { Client: Client7 } = require('es7')
const client = new Client7({ node: 'http://localhost:9200' })
const program = require('commander');

program.option('ping', 'ping server')

program.parse(process.argv)

if (program.ping) {
	client.cluster.health()
		.then(function(res){
			console.log(res.body)
		});
}
