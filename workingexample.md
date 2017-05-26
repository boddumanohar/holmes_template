## Analysis with Holmes - DEMO

#### Make sure that your service works fine.

If the docker-compose-command doesn't throw any other errors, you're well set up. You can verify that the services are running by executing "docker ps" (you may need to sudo). There you should see one line per service. If you look at the column "PORTS", you can see, what local port they are listening on. e.g. by default, the service zipmeta should be listening on 7770. Visiting http://127.0.0.1:7770 in your browser should display a help page for zipmeta.

As a first step, you should try, invoking the service directly without totem, in order to see, whether it works:
If you want to invoke zipmeta on your zip-file, place the zipfile in your /tmp-folder and then visit http://127.0.0.1:7770/analyze/?obj=myzip.zip in your browser (whereas you replace "myzip.zip" with the name of  your zip-file). If you get a JSON-answer that doesn't look like an error, you know the service is working and should try the same with totem.

#### Analysis with Totem

In order to do the same with totem, you first need to upload the sample to storage because that is where totem looks for it. You need to have Gateway and Storage running to be able to do so. If you're using Holmes-Toolbox to upload the sample (and it succeeds), it will return you a JSON-answer, which contains the SHA256-sum of the sample. Next just create a plain-text-file, where you put the SHA256-sum, the name of the sample (actually the name doesn't really matter), and the source (just choose "src1"), all space-separated, like the following:

```
d4748e7c9724c0d183a13e9582e7a8408892276821a8d03c50964010475b01e9 name src1
```

Then use Holmes-Toolbox to execute only zip-meta (replace myfile.txt by the path to the file you just created):

```
$ ./Holmes-Toolbox --gateway https://127.0.0.1:8090 --user test --pw test --tags '["tag1","tag2"]' --comment "mycomment" --insecure --tasking --file myfile.txt --tasks '{"ZIPMETA":[]}'
```


Holmes-Toolbox should report: "The server returned an empty string (success)". You will notice that Totem will print out quite a lot and you'll hope that you don't need to go through all of this. Both Gateways should print out quite a lot, as well. If everything went smoothly, Gateway should have put the tasking-request into Rabbitmq, Totem should have taken it from there, downloaded the corresponding sample from Storage, given it to ZIPMETA, collected the result, put the result back into RabbitMq, and Storage should have collected it from there and put it into the table "results".

You can verify that all of this worked by using cqlsh:

```
$ cqlsh -e "select * from holmes.results;"
```

This should give you the results-table and in the "results"-column you should be able to see the result of zipmeta (the same json-blob you saw earlier in your browser).
If you DON'T see this, something went wrong. You can try looking, whether your request went to the totem_misbehave-queue in RabbitMq (by using its web-interface). If your request isn't there either, it probably never reached totem and you probably have an error in your gateway-configuration.