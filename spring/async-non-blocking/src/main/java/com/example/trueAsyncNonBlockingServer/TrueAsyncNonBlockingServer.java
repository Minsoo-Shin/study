package com.example.trueAsyncNonBlockingServer;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;
import reactor.core.publisher.Mono;
import reactor.core.scheduler.Schedulers;

import java.time.Duration;

@SpringBootApplication
@RestController
public class TrueAsyncNonBlockingServer {

    public static void main(String[] args) {
        SpringApplication.run(TrueAsyncNonBlockingServer.class, args);
    }

    @GetMapping("/async-operation")
    public Mono<String> asyncOperation() {
        return Mono.fromCallable(() -> {
                    // 시간이 걸리는 작업 시뮬레이션
                    Thread.sleep(5000);
                    return "Async operation completed";
                })
                .subscribeOn(Schedulers.boundedElastic())
                .map(result -> {
                    System.out.println("Operation completed: " + result);
                    return result;
                })
                .onErrorReturn("An error occurred")
                .defaultIfEmpty("No result");
    }

    @GetMapping("/immediate-response")
    public Mono<String> immediateResponse() {
        Mono.fromCallable(() -> {
                    // 백그라운드에서 실행될 장시간 작업
                    Thread.sleep(5000);
                    System.out.println("Background task completed");
                    return "Background task result";
                })
                .subscribeOn(Schedulers.boundedElastic())
                .subscribe(); // 작업을 시작하고 결과를 기다리지 않음

        return Mono.just("Task started, check back later");
    }
}