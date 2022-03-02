<script lang="ts">
    export let label: string = "label"
    export let password: boolean = false

    let type: string = password ? "password" : "text"
    let value: string

    export function getValue() {
        return value
    }
</script>

<div class=" flex flex-col w-full gap-1">
    <span class=" text-sm">{label}</span>
    <!-- TODO: can we fix this?? -->
    {#if password}
        <input
            bind:value
            class=" bg-gray-300 rounded-lg p-2"
            type="password"
            required
        />
    {:else}
        <input
            bind:value
            class=" bg-gray-300 rounded-lg p-2"
            type="text"
            required
        />
    {/if}
</div>
