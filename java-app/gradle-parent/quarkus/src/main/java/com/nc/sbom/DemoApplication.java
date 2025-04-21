package com.nc.sbom;

import jakarta.ws.rs.GET;
import jakarta.ws.rs.Path;
import jakarta.ws.rs.core.Response;
import jakarta.ws.rs.core.Application;
import jakarta.ws.rs.ApplicationPath;

@ApplicationPath("/")
@Path("/")
public class DemoApplication extends Application {

    @SuppressWarnings("unused")
    @GET
    public Response sbom() {
        return Response.ok("Generating the SBOM!").build();
    }
}