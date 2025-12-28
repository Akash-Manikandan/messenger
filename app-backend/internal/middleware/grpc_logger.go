package middleware

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

type GRPCRequestLog struct {
	Timestamp  time.Time
	Method     string
	StatusCode string
	Latency    time.Duration
	IP         string
	Error      error
}

// UnaryServerInterceptor logs unary gRPC requests asynchronously
func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {
		start := time.Now()

		// Get client IP
		clientIP := "unknown"
		if p, ok := peer.FromContext(ctx); ok {
			clientIP = p.Addr.String()
		}

		// Capture request details
		reqLog := GRPCRequestLog{
			Timestamp: start,
			Method:    info.FullMethod,
			IP:        clientIP,
		}

		// Call the handler
		resp, err := handler(ctx, req)

		// Calculate response details
		reqLog.Latency = time.Since(start)
		reqLog.Error = err
		if err != nil {
			reqLog.StatusCode = status.Code(err).String()
		} else {
			reqLog.StatusCode = "OK"
		}

		// Log asynchronously
		go logGRPCRequest(reqLog)

		return resp, err
	}
}

// StreamServerInterceptor logs streaming gRPC requests asynchronously
func StreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(
		srv any,
		ss grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		start := time.Now()

		// Get client IP
		clientIP := "unknown"
		if p, ok := peer.FromContext(ss.Context()); ok {
			clientIP = p.Addr.String()
		}

		// Capture request details
		reqLog := GRPCRequestLog{
			Timestamp: start,
			Method:    info.FullMethod,
			IP:        clientIP,
		}

		// Call the handler
		err := handler(srv, ss)

		// Calculate response details
		reqLog.Latency = time.Since(start)
		reqLog.Error = err
		if err != nil {
			reqLog.StatusCode = status.Code(err).String()
		} else {
			reqLog.StatusCode = "OK"
		}

		// Log asynchronously
		go logGRPCRequest(reqLog)

		return err
	}
}

func logGRPCRequest(r GRPCRequestLog) {
	statusColor := colorGreen
	if r.Error != nil {
		statusColor = colorRed
	}
	latencyColor := getLatencyColor(r.Latency)

	fmt.Printf("%s[%s]%s %sgRPC%s %s%s%s | Status: %s%s%s | Duration: %s%v%s | IP: %s%s%s\n",
		colorCyan, r.Timestamp.Format("2006-01-02 15:04:05"), colorReset,
		colorPurple, colorReset,
		colorWhite, r.Method, colorReset,
		statusColor, r.StatusCode, colorReset,
		latencyColor, r.Latency, colorReset,
		colorBlue, r.IP, colorReset,
	)

	if r.Error != nil {
		fmt.Printf("  %sError: %v%s\n", colorRed, r.Error, colorReset)
	}
}
