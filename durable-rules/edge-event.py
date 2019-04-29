from durable.lang import *

with ruleset('motortoofastsignal'):
    @when_all((m.device == '562114e9e4b0385849b96cd8') & (m.parameter == 'RPM') & (m.value > 1200))
    def say_hello(c):
        print ('{0} warning triggered for high engine speed'.format(c.m.device))

    # @when_start
    # def start(host):
    #     host.post('motortoofastsignal', { 
    #       'device': '562114e9e4b0385849b96cd8',
    #       'parameter': 'RPM',
    #       'value': 1200 })

run_all()

# {
#   "name": "motortoofastsignal",
#   "condition": {
#     "device": "562114e9e4b0385849b96cd8",
#     "checks": [
#       {
#         "parameter": "RPM",
#         "operand1": "Integer.parseInt(value)",
#         "operation": ">",
#         "operand2": "1200"
#       }
#     ]
#   },
#   "action": {
#     "device": "56325f7ee4b05eaae5a89ce1",
#     "command": "56325f6de4b05eaae5a89cdc",
#     "body": "{\\\"value\\\":\\\"3\\\"}"
#   },
#   "log": "Patlite warning triggered for engine speed too high"
# }
