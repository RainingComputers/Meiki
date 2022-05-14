type VoidReturnAsyncFunc = (...args: any[]) => Promise<void>

export function debounce(
    func: VoidReturnAsyncFunc,
    threshold: number = 5000
): VoidReturnAsyncFunc {
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
