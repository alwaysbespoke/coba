package v1

type StatusCode int

// RFC 3261 Sections 25.1
const (
	// Informational
	Trying               StatusCode = 100
	Ringing              StatusCode = 180
	CallIsBeingForwarded StatusCode = 181
	Queued               StatusCode = 182
	SessionProgress      StatusCode = 183
	// Success
	OK StatusCode = 200
	// Redirection
	MultipleChoices    StatusCode = 300
	MovedPermanently   StatusCode = 301
	MovedTemporarily   StatusCode = 302
	UseProxy           StatusCode = 305
	AlternativeService StatusCode = 380
	// Client-Error
	BadRequest                       StatusCode = 400
	Unauthorized                     StatusCode = 401
	PaymentRequired                  StatusCode = 402
	Forbidden                        StatusCode = 403
	NotFound                         StatusCode = 404
	MethodNotAllowed                 StatusCode = 405
	NotAcceptable                    StatusCode = 406
	ProxyAuthenticationRequired      StatusCode = 407
	RequestTimeout                   StatusCode = 408
	Gone                             StatusCode = 410
	RequestEntityTooLarge            StatusCode = 413
	RequestURITooLarge               StatusCode = 414
	UnsupportedMediaType             StatusCode = 415
	UnsupportedURIScheme             StatusCode = 416
	BadExtension                     StatusCode = 420
	ExtensionRequired                StatusCode = 421
	IntervalTooBrief                 StatusCode = 423
	TemporarilyNotAvailable          StatusCode = 480
	CallLegOrTransactionDoesNotExist StatusCode = 481
	LoopDetected                     StatusCode = 482
	TooManyHops                      StatusCode = 483
	AddressIncomplete                StatusCode = 484
	Ambiguous                        StatusCode = 485
	BusyHere                         StatusCode = 486
	RequestTerminated                StatusCode = 487
	NotAcceptableHere                StatusCode = 488
	RequestPending                   StatusCode = 491
	Undecipherable                   StatusCode = 493
	// Server-Error
	InternalServerError    StatusCode = 500
	NotImplemented         StatusCode = 501
	BadGateway             StatusCode = 502
	ServiceUnavailable     StatusCode = 503
	ServerTimeOut          StatusCode = 504
	SIPVersionNotSupported StatusCode = 505
	MessageTooLarge        StatusCode = 513
	// Global-Failure
	BusyEverywhere       StatusCode = 600
	Decline              StatusCode = 603
	DoesNotExistAnywhere StatusCode = 604
	GlobalNotAcceptable  StatusCode = 606 // Not Acceptable
)
