# Blob Storage Server

A blob storage server can be useful in many scenarios - lets see if we can build one!

## The task 

The task is split into the several sections below with an expected time to code of a few hours (familiarity dependent). The 
majority of the work is in the first task and so the second half can be skipped if time is tight.


* Implement **either a durable or volatile** storage mechanism inline with the **`store/IStore`** specification
    * The solution should reside under `store/impl/durable|volatile`
    * The solution should **not** use an external service (DB|KVS)
    * The solution can utilise any public GitHub source
    * The solution should handle binary data < 100Mb
    * The solution should be adequately commented
    * The solution should include tests (unit/table or other)
    * Optional
        * Note any issues with this design  
        * Note any quick improvements that could be made
        * Note any other observations
    
    
* Implement a simple HTTP server to allow users to interact with our new storage (optional)
    * The solution should be contained within `main.go`
    * The solution can utilise any public GitHub source
    * The solution should attempt to implement the following routes
      * PUT /\<key\>/\<TTL:seconds\>
      * GET /\<key\>
      * DELETE /\<key\>

If you do not manage to complete any part of the task please feel free to fill in any gaps with useful comments which can be taken into
account upon review. 
