package com.nc.sbom;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.Bean;
import org.springframework.web.servlet.function.RouterFunction;
import org.springframework.web.servlet.function.RouterFunctions;
import org.springframework.web.servlet.function.ServerResponse;

@SpringBootApplication
public class DemoApplication {

    public static void main(String[] args) {
        SpringApplication.run(DemoApplication.class, args);
    }

    @SuppressWarnings("unused")
    @Bean
    RouterFunction<ServerResponse> routerFunctions() {
        return RouterFunctions.route()
                .GET("/", request -> ServerResponse.ok().body("Generating the SBOM!"))
                .build();
    }

}
