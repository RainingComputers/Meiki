// TODO: figure out typing

export function debounce(
    func: (...args: any[]) => Promise<any>,
    threshold: number = 5000
): any {
    let lastCall = new Date()

    return async function (...args: any[]) {
        const now = new Date()
        const delta = now.getTime() - lastCall.getTime()

        if (delta > threshold) {
            await func(...args)
            lastCall = now
        }
    }
}
