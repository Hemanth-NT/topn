// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Top N microservice pulls data from different internal service.It expose API's to search/aggregate operation on metrics collected from internal services.",
    "title": "Top N Micorservice",
    "version": "1.0.0"
  },
  "basePath": "/api/topn",
  "paths": {
    "/v1/analytics": {
      "post": {
        "description": "API to do statistical operation on metrics data.It includes getting average of some metrics over last x time interval. Similarly sum, mean, median etc of metrics based on some metric parameter. NOTE- statiscal operation can be done only on number or date fields on other fields operation fail with error.\nFor example:\n` + "`" + `` + "`" + `` + "`" + `\nPOST /api/topn/v1/analytics\n` + "`" + `` + "`" + `` + "`" + `\nSample request body will be :\n` + "`" + `` + "`" + `` + "`" + `\n  {\n    metrickey:\"CPU\",\n    startTime:12312312313,\n    endTime:12312312333,\n    pageSize:100,\n    pageNo:0,\n  }\n` + "`" + `` + "`" + `` + "`" + `\nSample response body will be :\n` + "`" + `` + "`" + `` + "`" + `\n\n{\n  metrickey:\"cpu\",\n  average: 12,\n  sum:2311,\n  metricCount: 122,\n  variance :5,\n  standardDeviation:2,\n  nextPage : {\n    metrickey:\"CPU\",\n    startTime:12312312313,\n    endTime:12312312333,\n    pageSize:100,\n    pageNo:1,\n  }\n}\n` + "`" + `` + "`" + `` + "`" + `\n",
        "tags": [
          "topn-microservice"
        ],
        "summary": "API to do statistical operation on metrics data.It includes getting average of some metrics over last x time interval. Similarly sum, mean, median etc of metrics based on some metric parameter. NOTE- statiscal operation can be done only on number or date fields on other fields operation fail with error.",
        "parameters": [
          {
            "description": "Metric Analytic request.",
            "name": "metricAnalyticRequest",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/metricAnalyticRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Metric request served successfully",
            "schema": {
              "$ref": "#/definitions/metricAnalyticResponse"
            }
          },
          "400": {
            "description": "Bad Request, Unable to parse the metric message.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/v1/gettopn": {
      "post": {
        "description": "For example metrics collected from service 1 are  below \n` + "`" + `` + "`" + `` + "`" + `\n[\n{\n \"customerAccount\": \"M Mobile\",\n \"eventTime\": \"Tue May 14 2019 18:11:15 GMT+0530 (India Standard Time)\"\n \"latitude\": \"123\",\n \"longitude\": \"123\",\n \"cell\": \"Bangalore India\"\n \"metricSource\": \"Internal Service A\",\n \"metricsParam\":{\n   \"cpuUsage\":\"90\"\n }  \n\n}\n\n{\n \"customerAccount\": \"P Mobile\",\n \"eventTime\": \"Tue May 14 2019 18:11:15 GMT+0530 (India Standard Time)\"\n \"latitude\": \"129\",\n \"longitude\": \"723\",\n \"cell\": \"Bangalore India\"\n \"metricSource\": \"Internal Service A\",\n \"metricsParam\":{\n   \"cpuUsage\":\"50\"\n }  \n\n}\n\n]\n\nAnd user want to know top 1 CPU consuming customer Account then it should result in.\n\ncustomerAccount= \"P Mobile\"\n` + "`" + `` + "`" + `` + "`" + `\n For example:\n ` + "`" + `` + "`" + `` + "`" + `\n POST /api/topn/v1/gettopn\n ` + "`" + `` + "`" + `` + "`" + `\n Sample request body will be :\n ` + "`" + `` + "`" + `` + "`" + `\n   {\n     metrickey:\"cpuUsage\",\n     filterCriteria:{\n       list:[{attributeName:\"\",\n         attributeValue:\"\",\n         operator:\"\"\n       }],\n       relational:[{attributeName:\"\",\n         attributeValue:\"\",\n         operator:\"\"\n       }],\n       rangeCriteria:[\n         {attributeName:\"\",\n         attributeValue:\"\",\n         operator:\"\"\n       }]\n     },\n     topValueCount:1000,\n     pageSize:100,\n     pageNo:1\n   }\n ` + "`" + `` + "`" + `` + "`" + `\n Sample response body will be :\n ` + "`" + `` + "`" + `` + "`" + `\n \n {\n   topNMetrics :[\n       {\n     metrickey:\"cpuUsage\",\n     metricValue:\"100\",\n   }...\n     ],\n   nextPage :{\n     metrickey :\"cpuUsage\",\n     filterCriteria:{},\n     topvaluecount:1000,\n     pagesize:100,\n     pageNo:2\n     \n   }\n }\n ` + "`" + `` + "`" + `` + "`" + `\n",
        "tags": [
          "topn-microservice"
        ],
        "summary": "To fetch TOP N information based on the statistical operation on metrics collected from internal services.",
        "parameters": [
          {
            "description": "Metric request.",
            "name": "metricRequest",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/topNRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Metric request served successfully",
            "schema": {
              "$ref": "#/definitions/topnresponse"
            }
          },
          "400": {
            "description": "Bad Request, Unable to parse the metric message.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/v1/pullmetrics": {
      "get": {
        "description": "All internal service will expose this endpoint. TOP N service will pull metrics from this API exposed by internal services.Reponse of this API will be array of metrics.\n\nFor example:\n` + "`" + `` + "`" + `` + "`" + `\nget /api/topn/v1/pushmetrics\n` + "`" + `` + "`" + `` + "`" + `\n\nmetric content:\n` + "`" + `` + "`" + `` + "`" + `\n  [{\n     \"customerAccount\": \"M Mobile\",\n     \"eventTime\": \"Tue May 14 2019 18:11:15 GMT+0530 (India Standard Time)\"\n     \"latitude\": \"123\",\n     \"longitude\": \"123\",\n     \"cell\": \"Bangalore India\"\n     \"metricSource\": \"a\",\n     \"metricsParam\":{\n        \"metricParam1\":\"a\",\n        \"metricParam2\":\"b\",\n        ...              \n      }\n  }]\n` + "`" + `` + "`" + `` + "`" + `\n",
        "produces": [
          "application/json"
        ],
        "tags": [
          "topn-microservice"
        ],
        "summary": "All Internal service will expose this API.TopN service will consume/pull metrics using this API.(NOTE THIS IS DOCUMENTS IN SAME SWAGGER FILE FOR DEMO PURPOSE BUT THIS SERVICE RUNS ON INTERNAL SERVICES NOT ON TOP N)",
        "responses": {
          "200": {
            "description": "Metrics pulled by TOP N service from other internal service.",
            "schema": {
              "$ref": "#/definitions/Metrics"
            }
          },
          "404": {
            "description": "Internal service not reachable.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal server error.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "Metric": {
      "description": "An metric feed",
      "type": "object",
      "required": [
        "metricSource",
        "customerAccount",
        "eventTime",
        "metricsParam",
        "latitude",
        "longitude",
        "cell"
      ],
      "properties": {
        "cell": {
          "description": "Humar readable location like Bangalore, Delhi etc",
          "type": "string"
        },
        "customerAccount": {
          "description": "service provider/customer id/Customer info used in analytical operation",
          "type": "string"
        },
        "eventTime": {
          "description": "Notification generation time",
          "type": "string"
        },
        "latitude": {
          "description": "latitude where device location",
          "type": "integer"
        },
        "longitude": {
          "description": "Longitude where device location",
          "type": "integer"
        },
        "metricSource": {
          "description": "Source or internal service from which metric is pulled",
          "type": "string"
        },
        "metricsParam": {
          "description": "JSON object it can have any mertics from device. Like usage of device, Frequent access of some site, CPU, Memory info etc But these information should adhare to standerd defined by TOP N Service",
          "type": "object"
        }
      }
    },
    "Metrics": {
      "type": "object",
      "required": [
        "metrics"
      ],
      "properties": {
        "metrics": {
          "description": "Array of Metric",
          "type": "array",
          "items": {
            "$ref": "#/definitions/Metric"
          }
        }
      }
    },
    "criteria": {
      "type": "object",
      "required": [
        "list",
        "range",
        "relational"
      ],
      "properties": {
        "list": {
          "description": "For example SQL IN condition query.",
          "type": "array",
          "items": {
            "$ref": "#/definitions/filterParam"
          }
        },
        "range": {
          "description": "For SQL Between condition call.",
          "type": "array",
          "items": {
            "$ref": "#/definitions/filterParam"
          }
        },
        "relational": {
          "description": "For SQL \u003e \u003c \u003c\u003e condition call.",
          "type": "array",
          "items": {
            "$ref": "#/definitions/filterParam"
          }
        }
      }
    },
    "filterParam": {
      "description": "Basic unit of filtering",
      "type": "object",
      "required": [
        "attributeValue",
        "attributeKey",
        "operator"
      ],
      "properties": {
        "attributeKey": {
          "description": "Value with which json field to be filtered",
          "type": "string"
        },
        "attributeValue": {
          "description": "Any JSON field to be filtered",
          "type": "string"
        },
        "operator": {
          "description": "operators \u003e ,\u003c ,= ,!=, etc",
          "type": "string"
        }
      }
    },
    "metricAnalyticRequest": {
      "type": "object",
      "required": [
        "metrickey",
        "startTime",
        "endTime",
        "pageSize",
        "pageNo"
      ],
      "properties": {
        "endTime": {
          "type": "string"
        },
        "metrickey": {
          "type": "string"
        },
        "pageNo": {
          "type": "integer"
        },
        "pageSize": {
          "type": "integer"
        },
        "startTime": {
          "type": "string"
        }
      }
    },
    "metricAnalyticResponse": {
      "type": "object",
      "required": [
        "metrickey",
        "average",
        "sum",
        "metricCount",
        "variance",
        "standardDeviation"
      ],
      "properties": {
        "average": {
          "type": "integer"
        },
        "metricCount": {
          "type": "integer"
        },
        "metrickey": {
          "type": "string"
        },
        "standardDeviation": {
          "type": "integer"
        },
        "sum": {
          "type": "integer"
        },
        "variance": {
          "type": "integer"
        }
      }
    },
    "topNRequest": {
      "description": "Model for TopN request API",
      "type": "object",
      "required": [
        "metrickey",
        "filterCriteria",
        "topvaluecount",
        "pagesize",
        "pageno"
      ],
      "properties": {
        "filterCriteria": {
          "type": "object",
          "$ref": "#/definitions/criteria"
        },
        "metrickey": {
          "description": "Json field on which ",
          "type": "string"
        },
        "pageno": {
          "description": "page number of response",
          "type": "integer"
        },
        "pagesize": {
          "description": "page size of response",
          "type": "integer"
        },
        "topvaluecount": {
          "description": "Top N count needed",
          "type": "integer"
        }
      }
    },
    "topResponseValue": {
      "type": "object",
      "required": [
        "metrickey",
        "metricValue"
      ],
      "properties": {
        "metricValue": {
          "description": "Values in sorted order",
          "type": "string"
        },
        "metrickey": {
          "description": "JSON field on which TOP N service filtered",
          "type": "string"
        }
      }
    },
    "topnresponse": {
      "type": "object",
      "properties": {
        "nextPage": {
          "type": "string"
        },
        "topNMetrics": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/topResponseValue"
          }
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Top N microservice pulls data from different internal service.It expose API's to search/aggregate operation on metrics collected from internal services.",
    "title": "Top N Micorservice",
    "version": "1.0.0"
  },
  "basePath": "/api/topn",
  "paths": {
    "/v1/analytics": {
      "post": {
        "description": "API to do statistical operation on metrics data.It includes getting average of some metrics over last x time interval. Similarly sum, mean, median etc of metrics based on some metric parameter. NOTE- statiscal operation can be done only on number or date fields on other fields operation fail with error.\nFor example:\n` + "`" + `` + "`" + `` + "`" + `\nPOST /api/topn/v1/analytics\n` + "`" + `` + "`" + `` + "`" + `\nSample request body will be :\n` + "`" + `` + "`" + `` + "`" + `\n  {\n    metrickey:\"CPU\",\n    startTime:12312312313,\n    endTime:12312312333,\n    pageSize:100,\n    pageNo:0,\n  }\n` + "`" + `` + "`" + `` + "`" + `\nSample response body will be :\n` + "`" + `` + "`" + `` + "`" + `\n\n{\n  metrickey:\"cpu\",\n  average: 12,\n  sum:2311,\n  metricCount: 122,\n  variance :5,\n  standardDeviation:2,\n  nextPage : {\n    metrickey:\"CPU\",\n    startTime:12312312313,\n    endTime:12312312333,\n    pageSize:100,\n    pageNo:1,\n  }\n}\n` + "`" + `` + "`" + `` + "`" + `\n",
        "tags": [
          "topn-microservice"
        ],
        "summary": "API to do statistical operation on metrics data.It includes getting average of some metrics over last x time interval. Similarly sum, mean, median etc of metrics based on some metric parameter. NOTE- statiscal operation can be done only on number or date fields on other fields operation fail with error.",
        "parameters": [
          {
            "description": "Metric Analytic request.",
            "name": "metricAnalyticRequest",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/metricAnalyticRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Metric request served successfully",
            "schema": {
              "$ref": "#/definitions/metricAnalyticResponse"
            }
          },
          "400": {
            "description": "Bad Request, Unable to parse the metric message.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/v1/gettopn": {
      "post": {
        "description": "For example metrics collected from service 1 are  below \n` + "`" + `` + "`" + `` + "`" + `\n[\n{\n \"customerAccount\": \"M Mobile\",\n \"eventTime\": \"Tue May 14 2019 18:11:15 GMT+0530 (India Standard Time)\"\n \"latitude\": \"123\",\n \"longitude\": \"123\",\n \"cell\": \"Bangalore India\"\n \"metricSource\": \"Internal Service A\",\n \"metricsParam\":{\n   \"cpuUsage\":\"90\"\n }  \n\n}\n\n{\n \"customerAccount\": \"P Mobile\",\n \"eventTime\": \"Tue May 14 2019 18:11:15 GMT+0530 (India Standard Time)\"\n \"latitude\": \"129\",\n \"longitude\": \"723\",\n \"cell\": \"Bangalore India\"\n \"metricSource\": \"Internal Service A\",\n \"metricsParam\":{\n   \"cpuUsage\":\"50\"\n }  \n\n}\n\n]\n\nAnd user want to know top 1 CPU consuming customer Account then it should result in.\n\ncustomerAccount= \"P Mobile\"\n` + "`" + `` + "`" + `` + "`" + `\n For example:\n ` + "`" + `` + "`" + `` + "`" + `\n POST /api/topn/v1/gettopn\n ` + "`" + `` + "`" + `` + "`" + `\n Sample request body will be :\n ` + "`" + `` + "`" + `` + "`" + `\n   {\n     metrickey:\"cpuUsage\",\n     filterCriteria:{\n       list:[{attributeName:\"\",\n         attributeValue:\"\",\n         operator:\"\"\n       }],\n       relational:[{attributeName:\"\",\n         attributeValue:\"\",\n         operator:\"\"\n       }],\n       rangeCriteria:[\n         {attributeName:\"\",\n         attributeValue:\"\",\n         operator:\"\"\n       }]\n     },\n     topValueCount:1000,\n     pageSize:100,\n     pageNo:1\n   }\n ` + "`" + `` + "`" + `` + "`" + `\n Sample response body will be :\n ` + "`" + `` + "`" + `` + "`" + `\n \n {\n   topNMetrics :[\n       {\n     metrickey:\"cpuUsage\",\n     metricValue:\"100\",\n   }...\n     ],\n   nextPage :{\n     metrickey :\"cpuUsage\",\n     filterCriteria:{},\n     topvaluecount:1000,\n     pagesize:100,\n     pageNo:2\n     \n   }\n }\n ` + "`" + `` + "`" + `` + "`" + `\n",
        "tags": [
          "topn-microservice"
        ],
        "summary": "To fetch TOP N information based on the statistical operation on metrics collected from internal services.",
        "parameters": [
          {
            "description": "Metric request.",
            "name": "metricRequest",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/topNRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Metric request served successfully",
            "schema": {
              "$ref": "#/definitions/topnresponse"
            }
          },
          "400": {
            "description": "Bad Request, Unable to parse the metric message.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/v1/pullmetrics": {
      "get": {
        "description": "All internal service will expose this endpoint. TOP N service will pull metrics from this API exposed by internal services.Reponse of this API will be array of metrics.\n\nFor example:\n` + "`" + `` + "`" + `` + "`" + `\nget /api/topn/v1/pushmetrics\n` + "`" + `` + "`" + `` + "`" + `\n\nmetric content:\n` + "`" + `` + "`" + `` + "`" + `\n  [{\n     \"customerAccount\": \"M Mobile\",\n     \"eventTime\": \"Tue May 14 2019 18:11:15 GMT+0530 (India Standard Time)\"\n     \"latitude\": \"123\",\n     \"longitude\": \"123\",\n     \"cell\": \"Bangalore India\"\n     \"metricSource\": \"a\",\n     \"metricsParam\":{\n        \"metricParam1\":\"a\",\n        \"metricParam2\":\"b\",\n        ...              \n      }\n  }]\n` + "`" + `` + "`" + `` + "`" + `\n",
        "produces": [
          "application/json"
        ],
        "tags": [
          "topn-microservice"
        ],
        "summary": "All Internal service will expose this API.TopN service will consume/pull metrics using this API.(NOTE THIS IS DOCUMENTS IN SAME SWAGGER FILE FOR DEMO PURPOSE BUT THIS SERVICE RUNS ON INTERNAL SERVICES NOT ON TOP N)",
        "responses": {
          "200": {
            "description": "Metrics pulled by TOP N service from other internal service.",
            "schema": {
              "$ref": "#/definitions/Metrics"
            }
          },
          "404": {
            "description": "Internal service not reachable.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal server error.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "Metric": {
      "description": "An metric feed",
      "type": "object",
      "required": [
        "metricSource",
        "customerAccount",
        "eventTime",
        "metricsParam",
        "latitude",
        "longitude",
        "cell"
      ],
      "properties": {
        "cell": {
          "description": "Humar readable location like Bangalore, Delhi etc",
          "type": "string"
        },
        "customerAccount": {
          "description": "service provider/customer id/Customer info used in analytical operation",
          "type": "string"
        },
        "eventTime": {
          "description": "Notification generation time",
          "type": "string"
        },
        "latitude": {
          "description": "latitude where device location",
          "type": "integer"
        },
        "longitude": {
          "description": "Longitude where device location",
          "type": "integer"
        },
        "metricSource": {
          "description": "Source or internal service from which metric is pulled",
          "type": "string"
        },
        "metricsParam": {
          "description": "JSON object it can have any mertics from device. Like usage of device, Frequent access of some site, CPU, Memory info etc But these information should adhare to standerd defined by TOP N Service",
          "type": "object"
        }
      }
    },
    "Metrics": {
      "type": "object",
      "required": [
        "metrics"
      ],
      "properties": {
        "metrics": {
          "description": "Array of Metric",
          "type": "array",
          "items": {
            "$ref": "#/definitions/Metric"
          }
        }
      }
    },
    "criteria": {
      "type": "object",
      "required": [
        "list",
        "range",
        "relational"
      ],
      "properties": {
        "list": {
          "description": "For example SQL IN condition query.",
          "type": "array",
          "items": {
            "$ref": "#/definitions/filterParam"
          }
        },
        "range": {
          "description": "For SQL Between condition call.",
          "type": "array",
          "items": {
            "$ref": "#/definitions/filterParam"
          }
        },
        "relational": {
          "description": "For SQL \u003e \u003c \u003c\u003e condition call.",
          "type": "array",
          "items": {
            "$ref": "#/definitions/filterParam"
          }
        }
      }
    },
    "filterParam": {
      "description": "Basic unit of filtering",
      "type": "object",
      "required": [
        "attributeValue",
        "attributeKey",
        "operator"
      ],
      "properties": {
        "attributeKey": {
          "description": "Value with which json field to be filtered",
          "type": "string"
        },
        "attributeValue": {
          "description": "Any JSON field to be filtered",
          "type": "string"
        },
        "operator": {
          "description": "operators \u003e ,\u003c ,= ,!=, etc",
          "type": "string"
        }
      }
    },
    "metricAnalyticRequest": {
      "type": "object",
      "required": [
        "metrickey",
        "startTime",
        "endTime",
        "pageSize",
        "pageNo"
      ],
      "properties": {
        "endTime": {
          "type": "string"
        },
        "metrickey": {
          "type": "string"
        },
        "pageNo": {
          "type": "integer"
        },
        "pageSize": {
          "type": "integer"
        },
        "startTime": {
          "type": "string"
        }
      }
    },
    "metricAnalyticResponse": {
      "type": "object",
      "required": [
        "metrickey",
        "average",
        "sum",
        "metricCount",
        "variance",
        "standardDeviation"
      ],
      "properties": {
        "average": {
          "type": "integer"
        },
        "metricCount": {
          "type": "integer"
        },
        "metrickey": {
          "type": "string"
        },
        "standardDeviation": {
          "type": "integer"
        },
        "sum": {
          "type": "integer"
        },
        "variance": {
          "type": "integer"
        }
      }
    },
    "topNRequest": {
      "description": "Model for TopN request API",
      "type": "object",
      "required": [
        "metrickey",
        "filterCriteria",
        "topvaluecount",
        "pagesize",
        "pageno"
      ],
      "properties": {
        "filterCriteria": {
          "type": "object",
          "$ref": "#/definitions/criteria"
        },
        "metrickey": {
          "description": "Json field on which ",
          "type": "string"
        },
        "pageno": {
          "description": "page number of response",
          "type": "integer"
        },
        "pagesize": {
          "description": "page size of response",
          "type": "integer"
        },
        "topvaluecount": {
          "description": "Top N count needed",
          "type": "integer"
        }
      }
    },
    "topResponseValue": {
      "type": "object",
      "required": [
        "metrickey",
        "metricValue"
      ],
      "properties": {
        "metricValue": {
          "description": "Values in sorted order",
          "type": "string"
        },
        "metrickey": {
          "description": "JSON field on which TOP N service filtered",
          "type": "string"
        }
      }
    },
    "topnresponse": {
      "type": "object",
      "properties": {
        "nextPage": {
          "type": "string"
        },
        "topNMetrics": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/topResponseValue"
          }
        }
      }
    }
  }
}`))
}
