import discord

intents = discord.Intents.default()
intents.message_content = True # Enable the message content intent

client = discord.Client(intents=intents)

@client.event
async def on_message(message):
    if isinstance(message.channel, discord.DMChannel):
        # This message was sent in a DM
        print(f"Received DM from {message.author}: {message.content}")
