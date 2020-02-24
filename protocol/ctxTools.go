package protocol

import "context"

const (
	ctxContactWriterClientKey = "ctx_contact_writer_client_key"
	ctxContactReaderClientKey = "ctx_contact_reader_client_key"
	ctxContactWriterServerKey = "ctx_contact_writer_server_key"
	ctxContactReaderServerKey = "ctx_contact_reader_server_key"
)

// ContactWriterClientFromContext returns a Contact Writer Client injected in Context
func ContactWriterClientFromContext(ctx context.Context) ContactWriterClient {
	cwI := ctx.Value(ctxContactWriterClientKey)

	if cw, ok := cwI.(ContactWriterClient); ok {
		return cw
	}

	return nil
}

// ContactReaderFromContext returns a Contact Reader Client injected in Context
func ContactReaderClientFromContext(ctx context.Context) ContactReaderClient {
	crI := ctx.Value(ctxContactReaderClientKey)

	if cr, ok := crI.(ContactReaderClient); ok {
		return cr
	}

	return nil
}

// ContactWriterServerFromContext returns a Contact Writer Server injected in Context
func ContactWriterServerFromContext(ctx context.Context) ContactWriterServer {
	cwI := ctx.Value(ctxContactWriterServerKey)

	if cw, ok := cwI.(ContactWriterServer); ok {
		return cw
	}

	return nil
}

// ContactReaderFromContext returns a Contact Reader Server injected in Context
func ContactReaderServerFromContext(ctx context.Context) ContactReaderServer {
	crI := ctx.Value(ctxContactReaderServerKey)

	if cr, ok := crI.(ContactReaderServer); ok {
		return cr
	}

	return nil
}

// InjectContactWriterClientInContext injects the contact writer client to the specified context
func InjectContactWriterClientInContext(ctx context.Context, cw ContactWriterClient) context.Context {
	return context.WithValue(ctx, ctxContactWriterClientKey, cw)
}

// InjectContactReaderClientInContext injects the contact reader client to the specified context
func InjectContactReaderClientInContext(ctx context.Context, cr ContactReaderClient) context.Context {
	return context.WithValue(ctx, ctxContactReaderClientKey, cr)
}

// InjectContactWriterServerInContext injects the contact writer client to the specified context
func InjectContactWriterServerInContext(ctx context.Context, cw ContactWriterServer) context.Context {
	return context.WithValue(ctx, ctxContactWriterServerKey, cw)
}

// InjectContactReaderServerInContext injects the contact reader client to the specified context
func InjectContactReaderServerInContext(ctx context.Context, cr ContactReaderServer) context.Context {
	return context.WithValue(ctx, ctxContactReaderServerKey, cr)
}
