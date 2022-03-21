<script lang="ts">
    export let showRender = false
    export let showEdit = true
</script>

<div class="flex flex-row flex-grow justify-center">
    {#if showEdit && !showRender}
        <div class=" bg-green-100 w-3/5">Something</div>
    {/if}
    {#if showRender && !showEdit}
        <div class=" bg-red-100 w-4/5">Something Else</div>
    {/if}
    {#if showRender && showEdit}
        <div class=" bg-green-100 flex-1">Something</div>
        <div class=" bg-red-100 flex-1">Something Else</div>
    {/if}
</div>
