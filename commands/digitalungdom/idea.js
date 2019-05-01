/* global db include guild */

const getUserByDiscordId = include( 'models/get' ).getUserByDiscordId;
const createNotificationEmbed = include( 'utils/embeds/createNotificationEmbed' );

module.exports = {
  name: 'idea',
  description: 'Skicka ett förslag till Digital Ungdom.',
  aliases: [ 'suggestion', 'förslag' ],
  group: 'digitalungdom',
  usage: 'idea <idea>',
  example: 'idea skriva kommentarer till er kod',
  serverOnly: false,
  adminOnly: false,
  async execute( message, args ) {
    if ( args.length === 0 ) return message.reply( 'Du måste skicka med ett kort meddelande.' );
    const idea = args.join( ' ' );
    const authorId = message.author.id;
    const authorUsername = message.author.username;
    let id = await getUserByDiscordId( authorId );
    if ( id ) id = id._id;

    await db.collection( 'notifications' ).insertOne( {
      'type': 'idea',
      'message': idea,
      'author': {
        'id': id,
        'discordId': authorId,
        'discordUsername': authorUsername
      }
    } );

    const notification = createNotificationEmbed( 'IDEA', idea, 65397, { 'id': authorId, 'name': authorUsername, 'url': message.author.displayAvatarURL } );
    const notificationChannel = guild.channels.find( ch => ch.name === 'notifications' );
    notificationChannel.send( '@here', { 'embed': notification } );

    return message.author.send( 'tack för din medverkan!' );
  },
};