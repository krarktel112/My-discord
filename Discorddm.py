import discord

@client.event
async def on_message(message):
    if isinstance(message.channel, discord.DMChannel):
        # This message was sent in a DM
        print(f"Received DM from {message.author}: {message.content}")
