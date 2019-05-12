Note:

 Straight lines indicate request.
 Dotted lines indicate response.
 Architecture is done considering application will be deployed on K8S.Architecture diagram also show flow. Each flow line is having identifier number explanation is provided for each number. 
 Authentication mechanism used OAuth2.
 Storage used for metrics Elastic search.
 

 
1) Client/Any device in network which need to connect to TOP N service will make HTTPS request to TOP N service with HTTP basic authentication.
2) Since Request dont have any bearer token request will not be served but user and password will be taken from HTTP header and user details passed to  KeyCloak identity and access management tool.
3) KeyCloak will check if credentials passed is available in its store (MySQL DB)(NOTE : THIRD PARTY ACCESS SERVER LIKE GOOGLE, FACEBOOK, GITHUB,CUSTOMER SPECIFIC CAN ALSO BE USED I AM CONSIDERING IT OPTIONAL HERE IT GIVES SCALING BENIFIT WE DONT HAVE TO WORRY ABOUT SCALING IF USER BASE GROWS).
4)MySQL returns credentials if its present.
5)If Credentials passed present in MySQL server keyCloak will prepate JWT token and share it with client/device.
6)Client with get JWT token which can be used for further communication.

 NOTE : POINT 1-6 IS REGISTRATION PROCESS
 
 
 
 
 Rest of the flow is is collection of Service registration, Metric collection, TOPN/Analytics API 