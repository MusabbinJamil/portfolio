const ANALYTICS_ENDPOINT = '/api/analytics';

interface TrackEvent {
	eventType: 'pageview' | 'click';
	target?: string;
	label?: string;
	referrer?: string;
}

function sendEvent(event: TrackEvent): void {
	const body = JSON.stringify(event);

	if (navigator.sendBeacon) {
		const blob = new Blob([body], { type: 'application/json' });
		navigator.sendBeacon(ANALYTICS_ENDPOINT, blob);
	} else {
		fetch(ANALYTICS_ENDPOINT, {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body,
			keepalive: true
		}).catch(() => {});
	}
}

export function trackPageView(): void {
	sendEvent({
		eventType: 'pageview',
		referrer: document.referrer || ''
	});
}

export function trackClick(target: string, label: string): void {
	sendEvent({
		eventType: 'click',
		target,
		label
	});
}
