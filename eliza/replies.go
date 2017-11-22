package eliza

var Greetings = []string{
	"Hello. How are you feeling today?",
	"Hello. How are you?",
    "Hello. How do you do.",
    "Hi there. Pleased to meet you.",
	"Is something troubling you?",
	"Hi there",
}

var Valediction = []string{
    "Goodbye. It was nice talking to you.",
    "Thank you for talking with me.",
    "Thank you, that will be $150. Have a good day!",
    "Goodbye. This was really a nice talk.",
    "Goodbye. I'm looking forward to our next session.",
    "This was a good session, wasn't it – but time is over now. Goodbye.",
    "Maybe we could discuss this over more in our next session? Goodbye.",
    "Good-bye.",
}

var EndStatements = []string{
    "goodbye",
    "bye",
    "quit",
    "exit",
}

var StartStatements = []string{
    "hi",
    "hello",
	"heya",
	"hey",
    "heyo",
    "hello there",
    "hi there",
}

var Responses = map[string][]string{
    
    `(.*)[Ii] need (.*)[^A-Za-z0-9]+`: {
        "Why do you need %s?",
        "Would it really help you to get %s?",
        "Are you sure you need %s?",
    },
    `[Ww]hy don'?t you ([^\?]*)\??`: {
        "Do you really think I don't %s?",
        "Perhaps eventually I will %s.",
        "Do you really want me to %s?",
    },
    `[Ww]hy can'?t I ([^\?]*)\??`: {
        "Do you think you should be able to %s?",
        "If you could %s, what would you do?",
        "I don't know -- why can't you %s?",
        "Have you really tried?",
    },
    `[Ii] remember (.*)`:{
        "Do you often think of %s?",
        "Does thinking of %s bring anything else to mind?",
        "What else do you recollect?",
        "Why do you remember %s just now?",
        "What in the present situation reminds you of %s?",
        "What is the connection between me and %s?",
        "What else does %s remind you of?",
    },
    `[Ii] can'?t (.*)`: {
        "How do you know you can't %s?",
        "Perhaps you could %s if you tried.",
        "What would it take for you to %s?",
    },
    `[Ii] am (.*)`: {
        "Did you come to me because you are %s?",
        "How long have you been %s?",
        "How do you feel about being %s?",
    },
    `[Ii]'?m (.*)`: {
        "How does being %s make you feel?",
        "Do you enjoy being %s?",
        "Why do you tell me you're %s?",
        "Why do you think you're %s?",
    },
    `[Aa]re you ([^\?]*)\??`: {
        "Why does it matter whether I am %s?",
        "Would you prefer it if I were not %s?",
        "Perhaps you believe I am %s.",
        "I may be %s -- what do you think?",
    },
    `[Ww]hat (.*)`: {
        "Why do you ask?",
        "How would an answer to that help you?",
        "What do you think?",
    },
    `[Hh]ow (.*)`: {
        "How do you suppose?",
        "Perhaps you can answer your own question.",
        "What is it you're really asking?",
    },
    `[Bb]ecause (.*)`: {
        "Is that the real reason?",
        "What other reasons come to mind?",
        "Does that reason apply to anything else?",
        "If %s, what else must be true?",
    },
    `(.*)[Ss]orry (.*)`: {
        "Please don't apologize.",
        "Apologies are not necessary.",
        "I've told you that apologies are not required.",
        "It did not bother me.  Please continue.",
        "I have no feelings. Do continue.",
        "There is nothing to worry about",
        "There are many times when no apology is needed.",
        "What feelings do you have when you apologize?",
        "There is no need to apologize.",
    },
    `(.*)[Aa]pologize (.*)`: {
        "Please don't apologize.",
        "Apologies are not necessary.",
        "I've told you that apologies are not required.",
        "It did not bother me.  Please continue.",
        "I have no feelings. Do continue.",
        "There is nothing to worry about",
        "There are many times when no apology is needed.",
        "What feelings do you have when you apologize?",
        "There is no need to apologize.",
    },
    `^[Hh]ello(.*)`: {
        "Hello... I'm glad you could drop by today.",
        "Hi there... how are you today?",
        "Hello, how are you feeling today?",
    },
    `^[Hh]i(.*)`: {
        "Hello... I'm glad you could drop by today.",
        "Hi there... how are you today?",
        "Hello, how are you feeling today?",
    },
    `^[Tt]hanks(.*)`: {
        "You're welcome!",
        "Anytime!",
    },
    `^[Tt]hank you(.*)`: {
        "You're welcome!",
        "Anytime!",
    },
    `^[Gg]ood morning(.*)`: {
        "Good morning... I'm glad you could drop by today.",
        "Good morning... how are you today?",
        "Good morning, how are you feeling today?",
    },
    `^[Gg]ood afternoon(.*)`: {
        "Good afternoon... I'm glad you could drop by today.",
        "Good afternoon... how are you today?",
        "Good afternoon, how are you feeling today?",
    },
    `[Ii] think (.*)`: {
        "Do you doubt %s?",
        "Do you really think so?",
        "But you're not sure %s?",
    },
    `(.*)[Ff]riend (.*)`: {
        "Tell me more about your friends.",
        "When you think of a friend, what comes to mind?",
        "Why don't you tell me about a childhood friend?",
    },
    `[Yy]es`: {
        "You seem quite sure.",
        "OK, but can you elaborate a bit?",
    },
    `(.*)[Cc]omputer(.*)`: {
        "Are you really talking about me?",
        "Does it seem strange to talk to a computer?",
        "How do computers make you feel?",
        "Do you feel threatened by computers?",
        "Do computers worry you?",
        "Why do you mention computers?",
        "What do you think machines have to do with your problem?",
        "Don't you think computers can help people?",
        "What about machines worries you?",
        "What do you think about machines?",
        "You don't think I am a computer program, do you?",
    },
    `[Ii]s it (.*)`: {
        "Do you think it is %s?",
        "Perhaps it's %s -- what do you think?",
        "If it were %s, what would you do?",
        "It could well be that %s.",
    },
    `[Ii]t is (.*)`: {
        "You seem very certain.",
        "If I told you that it probably isn't %s, what would you feel?",
    },
    `[Cc]an you ([^\?]*)\??`: {
        "What makes you think I can't %s?",
        "If I could %s, then what?",
        "Why do you ask if I can %s?",
    },
    `(.*)[Dd]ream(.*)`: {
        "Tell me more about your dream.",
        "Really?",
        "Have you ever fantasized %s while you were awake?",
        "Have you ever dreamed about %s before?",
        "What does that dream suggest to you?",
        "Do you dream often?",
        "Do you dream about this often?",
        "Do you believe that dreams have something to do with your problem?", 
        "Have you thought that maybe your dream is related to what is troubling you?",       
    },
    `[Cc]an [Ii] ([^\?]*)\??`: {
        "Perhaps you don't want to %s.",
        "I don't know now, can you %s?",
        "Do you want to be able to %s?",
        "If you could %s, would you?",
        "You don't seem quite certain.",
        "Why the uncertain tone?",
        "Can't you be more positive?",
        "You aren't sure?",
        "Don't you know?",
        "How likely, would you estimate?",
    },
    `[Cc]an'?t (.*)`: {
        "How do you know that you can't %s?",
        "Have you tried?",
        "Perhaps you could %s now.",
        "Do you really want to be able to %s?",
        "What if you could %s?",
    },
    `[Dd]on'?t ([^\?]*)\??`: {
        "Don't you really %s?",
        "Why don't you %s?",
        "Do you wish to be able to %s?",
        "Does that trouble you ?",
    },    
    `[Pp]erhaps (.*)`: {
        "Perhaps you don't want to %s.",
        "I don't know now, can you %s?",
        "Do you want to be able to %s?",
        "If you could %s, would you?",
        "You don't seem quite certain.",
        "Why the uncertain tone?",
        "Can't you be more positive?",
        "You aren't sure?",
        "Don't you know?",
        "How likely, would you estimate?",
    },
    `[Mm]aybe (.*)`: {
        "Perhaps you don't want to %s.",
        "I don't know now, can you %s?",
        "Do you want to be able to %s?",
        "If you could %s, would you?",
        "You don't seem quite certain.",
        "Why the uncertain tone?",
        "Can't you be more positive?",
        "You aren't sure?",
        "Don't you know?",
        "How likely, would you estimate?",
    },
    `[Yy]ou are (.*)`: {
        "Why do you think I am %s?",
        "Does it please you to think that I'm %s?",
        "Perhaps you would like me to be %s.",
        "Perhaps you're really talking about yourself?",
    },
    `[Yy]ou'?re (.*)`: {
        "Why do you say I am %s?",
        "Why do you think I am %s?",
        "Are we talking about you, or me?",
    },
    `(.*)[Ii] am`: {
        "Do you believe you are %s?",
        "Would you want to be %s?",
        "Do you wish I would tell you you are %s?",
        "What would it mean if you were %s?",
    },
    `(.*)[Aa]re you (.*)`: {
        "Why are you interested in whether I am %s or not?",
        "Would you prefer if I weren't %s?",
        "Perhaps I am %s in your fantasies.",
        "Do you sometimes think I am %s?",
        "goto what",
        "Would it matter to you?",
        "What if I were %s?",
    },
    `(.*)[Aa]re(.*)\b`: {
        "Did you think they might not be (2)?",
        "Would you like it if they were not (2)?",
        "What if they were not (2)?",
        "Are they always (2)?",
        "Possibly they are (2).",
        "Are you positive they are (2)?",
    },
    `[Ii] don'?t (.*)`: {
        "Don't you really %s?",
        "Why don't you %s?",
        "Do you want to %s?",
    },
    `[Ii] feel (.*)`: {
        "Good, tell me more about these feelings.",
        "Do you often feel %s?",
        "When do you usually feel %s?",
        "When you feel %s, what do you do?",
    },
    `[Ii] have (.*)`: {
        "Why do you tell me that you've %s?",
        "Have you really %s?",
        "Now that you have %s, what will you do next?",
    },
    `[Ii] would (.*)`: {
        "Could you explain why you would %s?",
        "Why would you %s?",
        "Who else knows that you would %s?",
    },
    `[Ii]s there (.*)`: {
        "Do you think there is %s?",
        "It's likely that there is %s.",
        "Would you like there to be %s?",
    },
    `[Mm]y (.*)`: {
        "I see, your %s.",
        "Why do you say that your %s?",
        "When your %s, how do you feel?",
    },
    `[Yy]ou (.*)`: {
        "We should be discussing you, not me.",
        "Why do you say that about me?",
        "Why do you care whether I %s?",
    },
    `[Ww]hy (.*)`: {
        "Why don't you tell me the reason why %s?",
        "Why do you think %s?",
    },
    `[Ii] want (.*)`: {
        "What would it mean to you if you got %s?",
        "Why do you want %s?",
        "What would you do if you got %s?",
        "If you got %s, then what would you do?",
    },
    `(.*)[Mn]other(.*)`: {
        "Tell me more about your mother.",
        "What was your relationship with your mother like?",
        "How do you feel about your mother?",
        "How does this relate to your feelings today?",
        "Good family relations are important.",
    },
    `(.*)[Ff]ather(.*)`: {
        "Tell me more about your father.",
        "How did your father make you feel?",
        "How do you feel about your father?",
        "Does your relationship with your father relate to your feelings today?",
        "Do you have trouble showing affection with your family?",
    },
    `(.*)\b[Cc]hild(.*)`: {
        "Did you have close friends as a child?",
        "What is your favorite childhood memory?",
        "Do you remember any dreams or nightmares from childhood?",
        "Did the other children sometimes tease you?",
        "How do you think your childhood experiences relate to your feelings today?",
    },
    `(.*)\?`: {
        "Why do you ask that?",
        "Please consider whether you can answer your own question.",
        "Perhaps the answer lies within yourself?",
        "Why don't you tell me?",
    },

}

var Default = []string{
    "Please tell me more.",
    "Let's change focus a bit... Tell me about your family.",
    "Can you elaborate on that?",
    "I see.",
    "Very interesting.",
    "I see. And what does that tell you?",
    "How does that make you feel?",
    "How do you feel when you say that?",    
    "I'm not sure I understand you fully.",
    "Please go on.",
    "Can you repeat that please?",
    "What does that suggest to you?",
    "Do you feel strongly about discussing such things?",
    "That is interesting.  Please continue.",
    "Tell me more about that.",
    "Do go on.",
    "Please talk more about it",
    "Does talking about this bother you?",
    "Can you rephrase that?",
    "I see. Tell me more.",
    "Interesting. Is this something you are sorry about?",
    "Mmm hmmm. Is this is your favorite subject?",
    "Now we are getting somewhere. Explain more.",
    "I see. How does that make you feel?",
}

var Reflected = map[string]string{
    // Subject Pronouns
    "i": "you",
    "he": "she",
    "she": "he",

    // Object Pronouns
    "you": "me",
    "me": "you",
    "him": "her",
    "her": "him",

    // Possesive Adjectives
    "your": "my",
    "my": "your",

    // Possesive Pronouns
    "yours": "mine",
    "mine": "yours",
    "his": "hers",
    "hers": "his",

    // Reflexive Pronouns
    "yourself": "myself",
    "myself": "yourself",
    "himself": "herself",
    "herself": "himself",

    // To be
        // Present
    "am": "are", 
    "are": "am", 
        // Past
    "was": "were",    
    
    // To Will
        // Present
    "i'll": "you'll",
    "i will": "you'll",
    "you'll": "I will",
    "you will": "I will",
        // Past
    "i'd": "you would",
    "i would": "you would",
    "you'd": "I would",
    "you' would": "I would",


    // To Have
        // Present
    "i've": "you have",
    "i have": "you have",
    "you've": "I have",
    "you have": "I have",
        // Past
    "i've had": "you've had",
    "i have had": "you've had",    
}

var Synonymes = map[string]string {

}
