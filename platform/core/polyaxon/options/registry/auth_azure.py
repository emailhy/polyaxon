from options import option_namespaces, option_subjects
from options.cache import FREQUENT_CACHE_TTL
from options.option import NAMESPACE_DB_OPTION_MARKER, Option, OptionStores
from options.types import CONF_TYPES

AUTH_AZURE_ENABLED = '{}{}{}'.format(option_namespaces.AUTH_AZURE,
                                     NAMESPACE_DB_OPTION_MARKER,
                                     option_subjects.ENABLED)
AUTH_AZURE_VERIFICATION_SCHEDULE = '{}{}{}'.format(option_namespaces.AUTH_AZURE,
                                                   NAMESPACE_DB_OPTION_MARKER,
                                                   option_subjects.VERIFICATION_SCHEDULE)
AUTH_AZURE_TENANT_ID = '{}{}{}'.format(option_namespaces.AUTH_AZURE,
                                       NAMESPACE_DB_OPTION_MARKER,
                                       option_subjects.TENANT_ID)
AUTH_AZURE_CLIENT_ID = '{}{}{}'.format(option_namespaces.AUTH_AZURE,
                                       NAMESPACE_DB_OPTION_MARKER,
                                       option_subjects.CLIENT_ID)
AUTH_AZURE_CLIENT_SECRET = '{}{}{}'.format(option_namespaces.AUTH_AZURE,  # noqa
                                           NAMESPACE_DB_OPTION_MARKER,
                                           option_subjects.CLIENT_SECRET)


class AuthAzureEnabled(Option):
    key = AUTH_AZURE_ENABLED
    is_global = True
    is_secret = False
    is_optional = True
    is_list = False
    typing = CONF_TYPES.BOOL
    store = OptionStores.DB_OPTION
    default = False
    options = None
    cache_ttl = FREQUENT_CACHE_TTL


class AuthAzureVerificationSchedule(Option):
    key = AUTH_AZURE_VERIFICATION_SCHEDULE
    is_global = True
    is_secret = False
    is_optional = True
    is_list = False
    typing = CONF_TYPES.INT
    store = OptionStores.DB_OPTION
    default = 0
    options = None
    cache_ttl = FREQUENT_CACHE_TTL


class AuthAzureTenantId(Option):
    key = AUTH_AZURE_TENANT_ID
    is_global = True
    is_secret = True
    is_optional = True
    is_list = False
    typing = CONF_TYPES.STR
    store = OptionStores.DB_OPTION
    default = None
    options = None
    cache_ttl = FREQUENT_CACHE_TTL


class AuthAzureClientId(Option):
    key = AUTH_AZURE_CLIENT_ID
    is_global = True
    is_secret = True
    is_optional = True
    is_list = False
    typing = CONF_TYPES.STR
    store = OptionStores.DB_OPTION
    default = None
    options = None
    cache_ttl = FREQUENT_CACHE_TTL


class AuthAzureClientSecret(Option):
    key = AUTH_AZURE_CLIENT_SECRET
    is_global = True
    is_secret = True
    is_optional = True
    is_list = False
    typing = CONF_TYPES.STR
    store = OptionStores.DB_OPTION
    default = None
    options = None
    cache_ttl = FREQUENT_CACHE_TTL
