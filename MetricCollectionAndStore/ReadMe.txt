MetricCollectionAndStore.png Shows architectural diagram and flow of collecting/pulling metrics from other internal service and storing it for further use cases.

NOTE: 
	Desging : 1)TOPN Service and intermediate service to process metrics are stateless.
			  2)KAFKA stream processing used because when number of internal service providing metrics increase and if processing of metrics is very intense on resource then we may not be able to scale Intermediate service to too many instance. In that case we cant even lose metrics so it can be kept in kafka stream and intermediate service can process data at its capacity(Tuned to best performance against resource provded).
			  3)Data/Metrics stored in elastic search reason is it provides faster search (response time requirnemnt). Elastic search scales easily by adding more nodes and changing secondary replicas. 
			  4)Each internal service have its own index. In N internal service are available there will be N index.
			  
I assume there is no need of aggregating metrics accross metrics collected from different internal service.			  

Details of flow is mentioned in the png image.